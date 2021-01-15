package models

import "github.com/wingfeng/idx/utils"

// UserClaims [...]
type UserClaims struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:Id;not null"`
	UserID       string `gorm:"index:IX_UserClaims_UserId;column:UserId;type:varchar(255);not null"`
	Users        User   `gorm:"foreignkey:UserId;"`
	ClaimType    string `gorm:"column:ClaimType;type:longtext"`
	ClaimValue   string `gorm:"column:ClaimValue;type:longtext"`
	utils.Record `gorm:"embedded"`
}
