package models

import "github.com/wingfeng/idx/utils"

// UserRoles [...]
type UserRoles struct {
	UserID       string `gorm:"primary_key;column:userid;type:varchar(255);not null"`
	Users        User   `gorm:"foreignkey:userid;"`
	RoleID       string `gorm:"primary_key;index:IX_UserRoles_RoleId;column:roleid;type:varchar(255);not null"`
	OUID         string `gorm:"primary_key;column:ouid;type:varchar(36);not null"`
	Roles        Role   `gorm:"foreignkey:roleid;reference:Id"`
	utils.Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *UserRoles) TableName() string {
// 	return "UserRoles"
// }

func (m *UserRoles) GetID() interface{} {
	return m
}
