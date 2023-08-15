package models

import "github.com/wingfeng/idx/utils"

// UserLogins [...]
type UserLogins struct {
	LoginProvider       string `gorm:"primary_key;column:loginprovider;type:varchar(255);not null"`
	ProviderKey         string `gorm:"primary_key;column:providerkey;type:varchar(255);not null"`
	ProviderDisplayName string `gorm:"column:providerdisplayname;type:text"`
	UserID              string `gorm:"index:IX_UserLogins_UserId;column:userid;type:varchar(255);not null"`
	Users               User   `gorm:"foreignkey:userid;"`
	utils.Record        `gorm:"embedded"`
}
