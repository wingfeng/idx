package models

import "github.com/wingfeng/idx/utils"

// UserTokens [...]
type UserTokens struct {
	UserID string `gorm:"primary_key;column:UserId;type:varchar(255);not null"`

	LoginProvider string `gorm:"primary_key;column:LoginProvider;type:varchar(255);not null"`
	Name          string `gorm:"primary_key;column:Name;type:varchar(255);not null"`
	Value         string `gorm:"column:Value;type:longtext"`
	Users         User   `gorm:"foreignkey:UserId"`
	utils.Record  `gorm:"embedded"`
}
