package models

import (
	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// Client [...]
type Client struct {
	ID                                int       `gorm:"primary_key;auto_Increment;column:id;not null"`
	Enabled                           bool      `gorm:"column:enabled;not null"`
	ClientCode                        string    `gorm:"column:clientcode;type:varchar(256);not null"`
	GrantTypes                        string    `gorm:"column:granttypes;type:varchar(256)"`
	Scopes                            string    `gorm:"column:scopes;type:varchar(256)"`
	Domains                           string    `gorm:"column:domains;type:varchar(1024)"`
	ProtocolType                      string    `gorm:"column:protocoltype;type:varchar(200);not null"`
	RequireClientSecret               bool      `gorm:"column:requireclientsecret;not null"`
	ClientName                        string    `gorm:"column:clientname;type:varchar(200)"`
	Description                       string    `gorm:"column:description;type:varchar(1000)"`
	ClientURI                         string    `gorm:"column:clienturi;type:varchar(2000)"`
	LogoURI                           string    `gorm:"column:logouri;type:varchar(2000)"`
	RequireConsent                    bool      `gorm:"column:requireconsent;not null"`
	AllowRememberConsent              bool      `gorm:"column:allowrememberconsent;not null"`
	AlwaysIncludeUserClaimsInIDToken  bool      `gorm:"column:alwaysincludeuserclaimsinidtoken;not null"`
	RequirePkce                       bool      `gorm:"column:requirepkce;not null"`
	AllowPlainTextPkce                bool      `gorm:"column:allowplaintextpkce;not null"`
	AllowAccessTokensViaBrowser       bool      `gorm:"column:allowaccesstokensviabrowser;not null"`
	FrontChannelLogoutURI             string    `gorm:"column:frontchannellogouturi;type:varchar(2000)"`
	FrontChannelLogoutSessionRequired bool      `gorm:"column:frontchannellogoutsessionrequired;not null"`
	BackChannelLogoutURI              string    `gorm:"column:backchannellogouturi;type:varchar(2000)"`
	BackChannelLogoutSessionRequired  bool      `gorm:"column:backchannellogoutsessionrequired;not null"`
	AllowOfflineAccess                bool      `gorm:"column:allowofflineaccess;not null"`
	IDentityTokenLifetime             int       `gorm:"column:identitytokenlifetime;type:int;not null"`
	AccessTokenLifetime               int       `gorm:"column:accesstokenlifetime;type:int;not null"`
	AuthorizationCodeLifetime         int       `gorm:"column:authorizationcodelifetime;type:int;not null"`
	ConsentLifetime                   int       `gorm:"column:consentlifetime;type:int"`
	AbsoluteRefreshTokenLifetime      int       `gorm:"column:absoluterefreshtokenlifetime;type:int;not null"`
	SlidingRefreshTokenLifetime       int       `gorm:"column:slidingrefreshtokenlifetime;type:int;not null"`
	RefreshTokenUsage                 int       `gorm:"column:refreshtokenusage;type:int;not null"`
	UpdateAccessTokenClaimsOnRefresh  bool      `gorm:"column:updateaccesstokenclaimsonrefresh;not null"`
	RefreshTokenExpiration            int       `gorm:"column:refreshtokenexpiration;type:int;not null"`
	AccessTokenType                   int       `gorm:"column:accesstokentype;type:int;not null"`
	EnableLocalLogin                  bool      `gorm:"column:enablelocallogin;not null"`
	IncludeJwtID                      bool      `gorm:"column:includejwtid;not null"`
	AlwaysSendClientClaims            bool      `gorm:"column:alwayssendclientclaims;not null"`
	ClientClaimsPrefix                string    `gorm:"column:clientclaimsprefix;type:varchar(200)"`
	PairWiseSubjectSalt               string    `gorm:"column:pairwisesubjectsalt;type:varchar(200)"`
	LastAccessed                      null.Time `gorm:"column:lastaccessed;"`
	UserSsoLifetime                   int       `gorm:"column:userssolifetime;type:int"`
	UserCodeType                      string    `gorm:"column:usercodetype;type:varchar(100)"`
	DeviceCodeLifetime                int       `gorm:"column:devicecodelifetime;type:int;not null"`
	NonEditable                       bool      `gorm:"column:noneditable;not null"`
	utils.Record                      `gorm:"embedded"`
}

func (c *Client) GetID() string {
	return c.ClientCode
}
func (c *Client) GetSecret() string {
	return "c.se"
}
func (c *Client) GetDomain() string {
	return c.Domains
}
func (c *Client) GetUserID() string {
	return ""
}
func (c *Client) GetRequireConsent() bool {
	return c.RequireConsent
}

func (c *Client) ValifyPassword(password string) bool {
	return c.RequireClientSecret
}
func (c *Client) GetRequireSecret() bool {
	return c.RequireClientSecret
}
