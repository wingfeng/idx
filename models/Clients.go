package models

import (
	"strings"
)

// Client [...]
type Client struct {
	Id                                int                            `gorm:"primary_key;auto_Increment;not null"`
	Enabled                           bool                           `gorm:"not null"`
	ClientId                          string                         `gorm:"uniqueIndex;type:varchar(256);not null"`
	GrantTypes                        string                         `gorm:"type:varchar(256)"`
	Scopes                            string                         `gorm:"type:varchar(256)"`
	RedirectUris                      string                         `gorm:"type:varchar(1024)"`
	ProtocolType                      string                         `gorm:"type:varchar(200);not null"`
	RequireSecret                     bool                           `gorm:"not null"`
	ClientName                        string                         `gorm:"type:varchar(200)"`
	Description                       string                         `gorm:"type:varchar(1000)"`
	ClientURI                         string                         `gorm:"type:varchar(2000)"`
	LogoURI                           string                         `gorm:"type:varchar(2000)"`
	RequireConsent                    bool                           `gorm:"not null"`
	AllowRememberConsent              bool                           `gorm:"not null"`
	AlwaysIncludeUserClaimsInIdToken  bool                           `gorm:"not null"`
	RequirePkce                       bool                           `gorm:"not null"`
	AllowPlainTextPkce                bool                           `gorm:"not null"`
	AllowAccessTokensViaBrowser       bool                           `gorm:"not null"`
	FrontChannelLogoutURI             string                         `gorm:"type:varchar(2000)"`
	FrontChannelLogoutSessionRequired bool                           `gorm:"not null"`
	BackChannelLogoutURI              string                         `gorm:"type:varchar(2000)"`
	BackChannelLogoutSessionRequired  bool                           `gorm:"not null"`
	AllowOfflineAccess                bool                           `gorm:"not null"`
	IdentityTokenLifetime             int                            `gorm:"type:int;not null"`
	AccessTokenLifetime               int                            `gorm:"type:int;not null"`
	AuthorizationCodeLifetime         int                            `gorm:"type:int;not null"`
	ConsentLifetime                   int                            `gorm:"type:int"`
	AbsoluteRefreshTokenLifetime      int                            `gorm:"type:int;not null"`
	SlidingRefreshTokenLifetime       int                            `gorm:"type:int;not null"`
	RefreshTokenUsage                 int                            `gorm:"type:int;not null"`
	UpdateAccessTokenClaimsOnRefresh  bool                           `gorm:"not null"`
	RefreshTokenExpiration            int                            `gorm:"type:int;not null"`
	EnableLocalLogin                  bool                           `gorm:"not null"`
	AlwaysSendClientClaims            bool                           `gorm:"not null"`
	ClientClaimsPrefix                string                         `gorm:"type:varchar(200)"`
	PairWiseSubjectSalt               string                         `gorm:"type:varchar(200)"`
	UserSsoLifetime                   int                            `gorm:"type:int"`
	UserCodeType                      string                         `gorm:"type:varchar(100)"`
	DeviceCodeLifetime                int                            `gorm:"type:int;not null"`
	LogoutUris                        []ClientPostLogoutRedirectURIs `gorm:"foreignKey:ClientId"`
	Secrets                           []ClientSecrets                `gorm:"foreignKey:ClientId"`
	Record                            `gorm:"embedded"`
}

func (c *Client) TableName() string {
	return "clients"
}
func (c *Client) GetScopes() []string {
	return strings.Split(c.Scopes, " ")
}
func (c *Client) GetClientId() string {
	return c.ClientId
}

func (c *Client) GetClientName() string {
	return c.ClientName
}
func (c *Client) GetGrantTypes() []string {
	return strings.Split(c.GrantTypes, " ")
}
func (c *Client) GetRequireConsent() bool {
	return c.RequireConsent
}
func (c *Client) GetRedirectUris() []string {
	return strings.Split(c.RedirectUris, " ")
}
func (c *Client) GetPostLogoutUris() []string {
	result := []string{}
	for _, v := range c.LogoutUris {
		result = append(result, v.PostLogoutRedirectURI)
	}
	return result
}

// GetSecret 方法用于获取当前客户端的密钥。
// 该方法没有参数。
// 返回值是一个字符串切片，包含当前客户端的密钥。
func (c *Client) GetSecret() []string {
	//Secret set to empty. then do not check the client secret.
	if !c.RequireSecret {
		return []string{}
	} else {
		result := []string{}
		for _, v := range c.Secrets {
			result = append(result, v.Value)
		}
		return result
	}
}
func (c *Client) GetRequirePKCE() bool {
	return c.RequirePkce
}
