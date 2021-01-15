package models

import "github.com/wingfeng/idx/utils"

// RoleClaims [...]
type RoleClaims struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:Id;not null"`
	RoleID       string `gorm:"index:IX_RoleClaims_RoleId;column:RoleId;type:varchar(255);not null"`
	Roles        Role   `gorm:"association_foreignkey:RoleId;foreignkey:Id"`
	ClaimType    string `gorm:"column:ClaimType;type:longtext"`
	ClaimValue   string `gorm:"column:ClaimValue;type:longtext"`
	utils.Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *RoleClaims) TableName() string {
// 	return "RoleClaims"
// }
