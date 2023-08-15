package models

import "github.com/wingfeng/idx/utils"

// UserClaims [...]
type UserClaims struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:id;not null"`
	UserID       string `gorm:"index:IX_UserClaims_UserId;column:userid;type:varchar(255);not null"`
	Users        User   `gorm:"foreignkey:userid;"`
	ClaimType    string `gorm:"column:claimtype;type:text"`
	ClaimValue   string `gorm:"column:claimvalue;type:text"`
	utils.Record `gorm:"embedded"`
}
