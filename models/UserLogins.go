package models

import "github.com/wingfeng/idx/utils"

// UserLogins [...]
type UserLogins struct {
	LoginProvider       string `gorm:"primary_key;column:LoginProvider;type:varchar(255);not null"`
	ProviderKey         string `gorm:"primary_key;column:ProviderKey;type:varchar(255);not null"`
	ProviderDisplayName string `gorm:"column:ProviderDisplayName;type:longtext"`
	UserID              string `gorm:"index:IX_UserLogins_UserId;column:UserId;type:varchar(255);not null"`
	Users               User   `gorm:"foreignkey:UserId;"`
	utils.Record        `gorm:"embedded"`
}
