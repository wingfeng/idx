package models

import "github.com/wingfeng/idx/utils"

// UserTokens [...]
type UserTokens struct {
	UserID string `gorm:"primary_key;column:userid;type:varchar(255);not null"`

	LoginProvider string `gorm:"primary_key;column:loginprovider;type:varchar(255);not null"`
	Name          string `gorm:"primary_key;column:name;type:varchar(255);not null"`
	Value         string `gorm:"column:value;type:text"`
	Users         User   `gorm:"foreignkey:userid"`
	utils.Record  `gorm:"embedded"`
}
