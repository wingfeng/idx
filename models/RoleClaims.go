package models

import "github.com/wingfeng/idx/utils"

// RoleClaims [...]
type RoleClaims struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:id;not null"`
	RoleID       string `gorm:"index:IX_RoleClaims_RoleId;column:roleid;type:varchar(255);not null"`
	Roles        Role   `gorm:"foreignkey:roleid;reference:Id"`
	ClaimType    string `gorm:"column:claimtype;type:text"`
	ClaimValue   string `gorm:"column:claimvalue;type:text"`
	utils.Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *RoleClaims) TableName() string {
// 	return "RoleClaims"
// }
