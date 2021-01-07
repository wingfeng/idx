package models

import "github.com/wingfeng/idx/utils"

// ClientRedirectURIs [...]
type ClientRedirectURIs struct {
	ID           int    `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	RedirectURI  string `gorm:"column:RedirectUri;type:varchar(2000);not null"`
	ClientID     int    `gorm:"index:IX_ClientRedirectUris_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Client `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}
