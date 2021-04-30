package models

import (
	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// Client [...]
type Client struct {
	ID                                int       `gorm:"primary_key;auto_Increment;column:Id;not null"`
	Enabled                           bool      `gorm:"column:Enabled;type:tinyint(1);not null"`
	ClientID                          string    `gorm:"unique;column:ClientId;type:varchar(200);not null"`
	GrantTypes                        string    `gorm:"column:GrantTypes;type:varchar(256)"`
	Scopes                            string    `gorm:"column:Scopes;type:varchar(256)"`
	Domains                           string    `gorm:"column:Domains;type:varchar(1024)"`
	ProtocolType                      string    `gorm:"column:ProtocolType;type:varchar(200);not null"`
	RequireClientSecret               bool      `gorm:"column:RequireClientSecret;type:tinyint(1);not null"`
	ClientName                        string    `gorm:"column:ClientName;type:varchar(200)"`
	Description                       string    `gorm:"column:Description;type:varchar(1000)"`
	ClientURI                         string    `gorm:"column:ClientUri;type:varchar(2000)"`
	LogoURI                           string    `gorm:"column:LogoUri;type:varchar(2000)"`
	RequireConsent                    bool      `gorm:"column:RequireConsent;type:tinyint(1);not null"`
	AllowRememberConsent              bool      `gorm:"column:AllowRememberConsent;type:tinyint(1);not null"`
	AlwaysIncludeUserClaimsInIDToken  bool      `gorm:"column:AlwaysIncludeUserClaimsInIdToken;type:tinyint(1);not null"`
	RequirePkce                       bool      `gorm:"column:RequirePkce;type:tinyint(1);not null"`
	AllowPlainTextPkce                bool      `gorm:"column:AllowPlainTextPkce;type:tinyint(1);not null"`
	AllowAccessTokensViaBrowser       bool      `gorm:"column:AllowAccessTokensViaBrowser;type:tinyint(1);not null"`
	FrontChannelLogoutURI             string    `gorm:"column:FrontChannelLogoutUri;type:varchar(2000)"`
	FrontChannelLogoutSessionRequired bool      `gorm:"column:FrontChannelLogoutSessionRequired;type:tinyint(1);not null"`
	BackChannelLogoutURI              string    `gorm:"column:BackChannelLogoutUri;type:varchar(2000)"`
	BackChannelLogoutSessionRequired  bool      `gorm:"column:BackChannelLogoutSessionRequired;type:tinyint(1);not null"`
	AllowOfflineAccess                bool      `gorm:"column:AllowOfflineAccess;type:tinyint(1);not null"`
	IDentityTokenLifetime             int       `gorm:"column:IdentityTokenLifetime;type:int(11);not null"`
	AccessTokenLifetime               int       `gorm:"column:AccessTokenLifetime;type:int(11);not null"`
	AuthorizationCodeLifetime         int       `gorm:"column:AuthorizationCodeLifetime;type:int(11);not null"`
	ConsentLifetime                   int       `gorm:"column:ConsentLifetime;type:int(11)"`
	AbsoluteRefreshTokenLifetime      int       `gorm:"column:AbsoluteRefreshTokenLifetime;type:int(11);not null"`
	SlidingRefreshTokenLifetime       int       `gorm:"column:SlidingRefreshTokenLifetime;type:int(11);not null"`
	RefreshTokenUsage                 int       `gorm:"column:RefreshTokenUsage;type:int(11);not null"`
	UpdateAccessTokenClaimsOnRefresh  bool      `gorm:"column:UpdateAccessTokenClaimsOnRefresh;type:tinyint(1);not null"`
	RefreshTokenExpiration            int       `gorm:"column:RefreshTokenExpiration;type:int(11);not null"`
	AccessTokenType                   int       `gorm:"column:AccessTokenType;type:int(11);not null"`
	EnableLocalLogin                  bool      `gorm:"column:EnableLocalLogin;type:tinyint(1);not null"`
	IncludeJwtID                      bool      `gorm:"column:IncludeJwtId;type:tinyint(1);not null"`
	AlwaysSendClientClaims            bool      `gorm:"column:AlwaysSendClientClaims;type:tinyint(1);not null"`
	ClientClaimsPrefix                string    `gorm:"column:ClientClaimsPrefix;type:varchar(200)"`
	PairWiseSubjectSalt               string    `gorm:"column:PairWiseSubjectSalt;type:varchar(200)"`
	LastAccessed                      null.Time `gorm:"column:LastAccessed;type:datetime(6)"`
	UserSsoLifetime                   int       `gorm:"column:UserSsoLifetime;type:int(11)"`
	UserCodeType                      string    `gorm:"column:UserCodeType;type:varchar(100)"`
	DeviceCodeLifetime                int       `gorm:"column:DeviceCodeLifetime;type:int(11);not null"`
	NonEditable                       bool      `gorm:"column:NonEditable;type:tinyint(1);not null"`
	utils.Record                      `gorm:"embedded"`
}

func (c *Client) GetID() string {
	return c.ClientID
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
