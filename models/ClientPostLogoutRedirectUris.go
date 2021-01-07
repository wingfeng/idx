package models

import "github.com/wingfeng/idx/utils"

// ClientPostLogoutRedirectURIs [...]
type ClientPostLogoutRedirectURIs struct {
	ID                    int    `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	PostLogoutRedirectURI string `gorm:"column:PostLogoutRedirectUri;type:varchar(2000);not null"`
	ClientID              int    `gorm:"index:IX_ClientPostLogoutRedirectUris_ClientId;column:ClientId;type:int(11);not null"`
	Clients               Client `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	utils.Record          `gorm:"embedded"`
}
