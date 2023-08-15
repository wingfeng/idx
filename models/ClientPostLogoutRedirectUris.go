package models

import "github.com/wingfeng/idx/utils"

// ClientPostLogoutRedirectURIs [...]
type ClientPostLogoutRedirectURIs struct {
	ID                    int    `gorm:"primary_key;auto_Increment;column:id;not null"`
	PostLogoutRedirectURI string `gorm:"column:postlogoutredirecturi;type:varchar(2000);not null"`
	ClientID              int    `gorm:"index:IX_ClientPostLogoutRedirectUris_ClientId;column:clientid;type:int;not null"`
	Clients               Client `gorm:"foreignKey:clientid;references:id"`
	utils.Record          `gorm:"embedded"`
}
