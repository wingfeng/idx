package models

import "github.com/wingfeng/idx/utils"

// UserClaims [...]
type UserClaims struct {
	ID     int    `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	UserID string `gorm:"index:IX_UserClaims_UserId;column:UserId;type:varchar(255);not null"`
	//	Users        User   `gorm:"association_foreignkey:UserId;foreignkey:Id"`
	ClaimType    string `gorm:"column:ClaimType;type:longtext"`
	ClaimValue   string `gorm:"column:ClaimValue;type:longtext"`
	utils.Record `gorm:"embedded"`
}
