package core

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	log "github.com/cihub/seelog"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/store"
)

type Manager struct {
	HTTPScheme        string
	codeExp           time.Duration
	gtcfg             map[oauth2.GrantType]*manage.Config
	rcfg              *manage.RefreshingConfig
	authorizeGenerate oauth2.AuthorizeGenerate
	accessGenerate    oauth2.AccessGenerate
	tokenStore        oauth2.TokenStore
	clientStore       store.ClientStore
}

func NewDefaultManager() *Manager {
	m := NewManager()
	m.HTTPScheme = "http"
	// default implementation
	m.authorizeGenerate = generates.NewAuthorizeGenerate()
	m.accessGenerate = generates.NewAccessGenerate()

	return m
}
func NewManager() *Manager {
	return &Manager{
		gtcfg: make(map[oauth2.GrantType]*manage.Config),
	}
}
func (m *Manager) SetTokenStore(tStore oauth2.TokenStore) {
	m.tokenStore = tStore
}
func (m *Manager) SetClientStore(clientStore *store.ClientStore) {
	m.clientStore = *clientStore
}
func (m *Manager) MapAccessGenerate(gen oauth2.AccessGenerate) {
	m.accessGenerate = gen
}
func (m *Manager) GetClient(ctx context.Context, clientID string) (cli oauth2.ClientInfo, err error) {
	cli, err = m.clientStore.GetByID(ctx, clientID)
	if err != nil {
		return
	} else if cli == nil {
		err = errors.ErrInvalidClient
	}
	return
}

//GenerateAuthToken generate the authorization token(code)
func (m *Manager) GenerateAuthToken(ctx context.Context, rt oauth2.ResponseType, tgr *oauth2.TokenGenerateRequest) (authToken oauth2.TokenInfo, err error) {
	cli, err := m.clientStore.GetByID(ctx, tgr.ClientID)
	if err != nil {
		return nil, err
	} else if tgr.RedirectURI != "" {
		if err := m.validateURI(cli, tgr.RedirectURI); err != nil {
			return nil, err
		}
	}

	ti := models.NewToken()
	ti.SetClientID(tgr.ClientID)
	ti.SetUserID(tgr.UserID)
	ti.SetRedirectURI(tgr.RedirectURI)
	ti.SetScope(tgr.Scope)
	ti.SetState(tgr.State)
	ti.SetNonce(tgr.Nonce)
	iss := fmt.Sprintf("%s://%s", m.HTTPScheme, tgr.Request.Host)
	ti.SetIssuer(iss)

	createAt := time.Now()

	td := &oauth2.GenerateBasic{
		Client:    cli,
		UserID:    tgr.UserID,
		CreateAt:  createAt,
		TokenInfo: ti,
		Request:   tgr.Request,
	}
	srt := oauth2.ResponseType(strings.Fields(string(rt))[0])
	switch srt {
	case oauth2.Code:
		codeExp := m.codeExp
		if codeExp == 0 {
			codeExp = manage.DefaultCodeExp
		}
		ti.SetCodeCreateAt(createAt)
		ti.SetCodeExpiresIn(codeExp)
		if exp := tgr.AccessTokenExp; exp > 0 {
			ti.SetAccessExpiresIn(exp)
		}
		if tgr.CodeChallenge != "" {
			ti.SetCodeChallenge(tgr.CodeChallenge)
			ti.SetCodeChallengeMethod(tgr.CodeChallengeMethod)
		}

		tv, err := m.authorizeGenerate.Token(ctx, td)
		if err != nil {
			return nil, err
		}
		ti.SetCode(tv)
	case oauth2.Token, oauth2.IDToken:
		// set access token expires
		icfg := m.grantConfig(oauth2.Implicit)
		aexp := icfg.AccessTokenExp
		if exp := tgr.AccessTokenExp; exp > 0 {
			aexp = exp
		}
		ti.SetAccessCreateAt(createAt)
		ti.SetAccessExpiresIn(aexp)

		if icfg.IsGenerateRefresh {
			ti.SetRefreshCreateAt(createAt)
			ti.SetRefreshExpiresIn(icfg.RefreshTokenExp)
		}

		tv, rv, err := m.accessGenerate.Token(ctx, td, icfg.IsGenerateRefresh)
		if err != nil {
			return nil, err
		}
		ti.SetAccess(tv)

		if rv != "" {
			ti.SetRefresh(rv)
		}
	}

	err = m.tokenStore.Create(ctx, ti)
	if err != nil {
		return nil, err
	}
	return ti, nil
}

// get authorization code data
func (m *Manager) getAuthorizationCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	ti, err := m.tokenStore.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	} else if ti == nil || ti.GetCode() != code || ti.GetCodeCreateAt().Add(ti.GetCodeExpiresIn()).Before(time.Now()) {
		err = errors.ErrInvalidAuthorizeCode
		return nil, errors.ErrInvalidAuthorizeCode
	}
	return ti, nil
}

// delete authorization code data
func (m *Manager) delAuthorizationCode(ctx context.Context, code string) error {
	return m.tokenStore.RemoveByCode(ctx, code)
}

// get and delete authorization code data
func (m *Manager) getAndDelAuthorizationCode(ctx context.Context, tgr *oauth2.TokenGenerateRequest) (oauth2.TokenInfo, error) {
	code := tgr.Code
	ti, err := m.getAuthorizationCode(ctx, code)
	if err != nil {
		return nil, err
	} else if ti.GetClientID() != tgr.ClientID {
		return nil, errors.ErrInvalidAuthorizeCode
	} else if codeURI := ti.GetRedirectURI(); codeURI != "" && codeURI != tgr.RedirectURI {
		return nil, errors.ErrInvalidAuthorizeCode
	}

	err = m.delAuthorizationCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return ti, nil
}

func (m *Manager) validateCodeChallenge(ti oauth2.TokenInfo, ver string) error {
	cc := ti.GetCodeChallenge()
	// early return
	if cc == "" && ver == "" {
		return nil
	}
	if cc == "" {
		return errors.ErrMissingCodeVerifier
	}
	if ver == "" {
		return errors.ErrMissingCodeVerifier
	}
	ccm := ti.GetCodeChallengeMethod()
	if ccm.String() == "" {
		ccm = oauth2.CodeChallengePlain
	}
	if !ccm.Validate(cc, ver) {
		return errors.ErrInvalidCodeChallenge
	}
	return nil
}

//GenerateAccessToken generate the access token
func (m *Manager) GenerateAccessToken(ctx context.Context, gt oauth2.GrantType, tgr *oauth2.TokenGenerateRequest) (accessToken oauth2.TokenInfo, err error) {
	cli, err := m.clientStore.GetByID(ctx, tgr.ClientID)
	if err != nil {
		return nil, err
	}
	// if cliPass, ok := cli.(oauth2.ClientPasswordVerifier); ok {
	// 	if !cliPass.VerifyPassword(tgr.ClientSecret) {
	// 		return nil, errors.ErrInvalidClient
	// 	}
	// } else
	if m.clientStore.ValidateSecret(tgr.ClientID, tgr.ClientSecret) != nil {
		return nil, errors.ErrInvalidClient
	}
	if tgr.RedirectURI != "" {
		if err := m.validateURI(cli, tgr.RedirectURI); err != nil {
			return nil, err
		}
	}

	if gt == oauth2.AuthorizationCode {
		ti, err := m.getAndDelAuthorizationCode(ctx, tgr)
		if err != nil {
			return nil, err
		}
		if err := m.validateCodeChallenge(ti, tgr.CodeVerifier); err != nil {
			return nil, err
		}
		tgr.UserID = ti.GetUserID()
		tgr.Scope = ti.GetScope()
		tgr.State = ti.GetState()
		if exp := ti.GetAccessExpiresIn(); exp > 0 {
			tgr.AccessTokenExp = exp
		}
	}

	ti := models.NewToken()
	ti.SetClientID(tgr.ClientID)
	ti.SetUserID(tgr.UserID)
	ti.SetRedirectURI(tgr.RedirectURI)
	ti.SetScope(tgr.Scope)
	ti.SetState(tgr.State)

	iss := fmt.Sprintf("%s://%s", m.HTTPScheme, tgr.Request.Host)
	ti.SetIssuer(iss)
	createAt := time.Now()
	ti.SetAccessCreateAt(createAt)

	// set access token expires
	gcfg := m.grantConfig(gt)
	aexp := gcfg.AccessTokenExp
	if exp := tgr.AccessTokenExp; exp > 0 {
		aexp = exp
	}
	ti.SetAccessExpiresIn(aexp)
	if gcfg.IsGenerateRefresh {
		ti.SetRefreshCreateAt(createAt)
		ti.SetRefreshExpiresIn(gcfg.RefreshTokenExp)
	}

	td := &oauth2.GenerateBasic{
		Client:    cli,
		UserID:    tgr.UserID,
		CreateAt:  createAt,
		TokenInfo: ti,
		Request:   tgr.Request,
	}

	av, rv, err := m.accessGenerate.Token(ctx, td, gcfg.IsGenerateRefresh)
	if err != nil {
		return nil, err
	}
	ti.SetAccess(av)

	if rv != "" {
		ti.SetRefresh(rv)
	}

	err = m.tokenStore.Create(ctx, ti)
	if err != nil {
		return nil, err
	}

	return ti, nil

}

// refreshing an access token
func (m *Manager) RefreshAccessToken(ctx context.Context, tgr *oauth2.TokenGenerateRequest) (accessToken oauth2.TokenInfo, err error) {
	cli, err := m.GetClient(ctx, tgr.ClientID)
	if err != nil {
		return nil, err
	} else if tgr.ClientSecret != cli.GetSecret() {
		return nil, errors.ErrInvalidClient
	}

	ti, err := m.LoadRefreshToken(ctx, tgr.Refresh)
	if err != nil {
		return nil, err
	} else if ti.GetClientID() != tgr.ClientID {
		return nil, errors.ErrInvalidRefreshToken
	}

	oldAccess, oldRefresh := ti.GetAccess(), ti.GetRefresh()

	td := &oauth2.GenerateBasic{
		Client:    cli,
		UserID:    ti.GetUserID(),
		CreateAt:  time.Now(),
		TokenInfo: ti,
		Request:   tgr.Request,
	}

	rcfg := manage.DefaultRefreshTokenCfg
	if v := m.rcfg; v != nil {
		rcfg = v
	}

	ti.SetAccessCreateAt(td.CreateAt)
	if v := rcfg.AccessTokenExp; v > 0 {
		ti.SetAccessExpiresIn(v)
	}

	if v := rcfg.RefreshTokenExp; v > 0 {
		ti.SetRefreshExpiresIn(v)
	}

	if rcfg.IsResetRefreshTime {
		ti.SetRefreshCreateAt(td.CreateAt)
	}

	if scope := tgr.Scope; scope != "" {
		ti.SetScope(scope)
	}

	tv, rv, err := m.accessGenerate.Token(ctx, td, rcfg.IsGenerateRefresh)
	if err != nil {
		return nil, err
	}

	ti.SetAccess(tv)
	if rv != "" {
		ti.SetRefresh(rv)
	}

	if err := m.tokenStore.Create(ctx, ti); err != nil {
		return nil, err
	}

	if rcfg.IsRemoveAccess {
		// remove the old access token
		if err := m.tokenStore.RemoveByAccess(ctx, oldAccess); err != nil {
			return nil, err
		}
	}

	if rcfg.IsRemoveRefreshing && rv != "" {
		// remove the old refresh token
		if err := m.tokenStore.RemoveByRefresh(ctx, oldRefresh); err != nil {
			return nil, err
		}
	}

	if rv == "" {
		ti.SetRefresh("")
		ti.SetRefreshCreateAt(time.Now())
		ti.SetRefreshExpiresIn(0)
	}

	return ti, nil
}

// use the access token to delete the token information
func (m *Manager) RemoveAccessToken(ctx context.Context, access string) (err error) {
	if access == "" {
		return errors.ErrInvalidAccessToken
	}
	return m.tokenStore.RemoveByAccess(ctx, access)
}

// use the refresh token to delete the token information
func (m *Manager) RemoveRefreshToken(ctx context.Context, refresh string) (err error) {
	if refresh == "" {
		return errors.ErrInvalidAccessToken
	}
	return m.tokenStore.RemoveByRefresh(ctx, refresh)
}

// according to the access token for corresponding token information
func (m *Manager) LoadAccessToken(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	if access == "" {
		return nil, errors.ErrInvalidAccessToken
	}

	ct := time.Now()
	ti, err := m.tokenStore.GetByAccess(ctx, access)
	if err != nil {
		return nil, err
	} else if ti == nil || ti.GetAccess() != access {
		return nil, errors.ErrInvalidAccessToken
	} else if ti.GetRefresh() != "" && ti.GetRefreshExpiresIn() != 0 &&
		ti.GetRefreshCreateAt().Add(ti.GetRefreshExpiresIn()).Before(ct) {
		return nil, errors.ErrExpiredRefreshToken
	} else if ti.GetAccessExpiresIn() != 0 &&
		ti.GetAccessCreateAt().Add(ti.GetAccessExpiresIn()).Before(ct) {
		return nil, errors.ErrExpiredAccessToken
	}
	return ti, nil
}

// according to the refresh token for corresponding token information
func (m *Manager) LoadRefreshToken(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	if refresh == "" {
		return nil, errors.ErrInvalidRefreshToken
	}

	ti, err := m.tokenStore.GetByRefresh(ctx, refresh)
	if err != nil {
		return nil, err
	} else if ti == nil || ti.GetRefresh() != refresh {
		return nil, errors.ErrInvalidRefreshToken
	} else if ti.GetRefreshExpiresIn() != 0 && // refresh token set to not expire
		ti.GetRefreshCreateAt().Add(ti.GetRefreshExpiresIn()).Before(time.Now()) {
		return nil, errors.ErrExpiredRefreshToken
	}
	return ti, nil
}
func (m *Manager) validateURI(cli *idxmodels.Client, rawuri string) error {
	url, err := url.Parse(rawuri)
	if err != nil {
		log.Errorf("传入URL错误!%s", url)
	}

	allowUris, err := m.clientStore.GetClientRedirectUris(cli.ID)
	if err != nil {
		log.Error("校验返回Uri错误!")
	}
	for _, s := range allowUris {

		surl, _ := url.Parse(s)
		//当scheme和host都相同就确认可以返回。
		if strings.EqualFold(url.Scheme, surl.Scheme) && strings.EqualFold(url.Host, surl.Host) {
			return nil
		}
	}
	return log.Errorf("不合法的Uri:%s", rawuri)
}

// get grant type config
func (m *Manager) grantConfig(gt oauth2.GrantType) *manage.Config {
	if c, ok := m.gtcfg[gt]; ok && c != nil {
		return c
	}
	switch gt {
	case oauth2.AuthorizationCode:
		return manage.DefaultAuthorizeCodeTokenCfg
	case oauth2.Implicit:
		return manage.DefaultImplicitTokenCfg
	case oauth2.PasswordCredentials:
		return manage.DefaultPasswordTokenCfg
	case oauth2.ClientCredentials:
		return manage.DefaultClientTokenCfg
	}
	return &manage.Config{}
}
