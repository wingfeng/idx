package core

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/cihub/seelog"
	"github.com/go-session/session"
	"github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
)

type OpenIDExtend struct {
	PrivateKeyByets []byte
	UserStore       *store.DbUserStore
	ClientStore     *store.ClientStore
}

func NewOpenIDExtend() *OpenIDExtend {
	ext := &OpenIDExtend{}
	return ext
}

// func (oidext *OpenIDExtend) ClientScopeHandler(clientid, scope string) (allow bool, err error) {
// 	scopes := strings.Split(scope, " ")
// 	supportScopes := oidext.ClientStore.GetClientScopes(clientid)
// 	for _, s := range scopes {
// 		isSupport := false
// 		for _, ss := range supportScopes {
// 			if strings.EqualFold(s, ss) {
// 				isSupport = true
// 				break
// 			}
// 		}
// 		if !isSupport {
// 			return false, fmt.Errorf("Client:%s Scope:%s not Supported", clientid, s)
// 		}
// 	}
// 	log.Debugf("Validate Client %s Scope:%s", clientid, scope)
// 	return true, nil
// }
func (oidext *OpenIDExtend) UserAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}
		returnURI := r.Form
		state := r.Form.Get("state")
		if !strings.EqualFold(state, "") {
			store.Set("state", state)
		}
		store.Set("ReturnUri", returnURI)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	//	store.Delete("LoggedInUserID")
	store.Save()
	return
}
func (oidext *OpenIDExtend) PasswordAuthorizationHandler(username, password string) (userID string, err error) {
	user, err := oidext.UserStore.GetUserByAccount(username)

	if err != nil {
		log.Errorf("获取用户%s信息错误,Error:%s", username, err.Error())
		return "", err
	}

	pwdVerified, err := utils.VerifyPassword(user.PasswordHash, password)
	if err != nil {
		log.Errorf("校验%s用户密码信息错误,Error:%s", username, err.Error())
		return "", err
	}
	if pwdVerified {
		return user.ID, nil
	}

	return "", fmt.Errorf("用户%s密码信息错误!", username)

}
