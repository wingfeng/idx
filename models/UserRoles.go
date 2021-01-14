package models

import "github.com/wingfeng/idx/utils"

// UserRoles [...]
type UserRoles struct {
	UserID string `gorm:"primary_key;column:UserId;type:varchar(255);not null"`
	//	Users        User   `gorm:"association_foreignkey:UserId;foreignkey:Id"`
	RoleID string `gorm:"primary_key;index:IX_UserRoles_RoleId;column:RoleId;type:varchar(255);not null"`
	OUID   string `gorm:"primary_key;column:OUId;type:varchar(36);not null"`
	//Roles        Role   `gorm:"association_foreignkey:RoleId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *UserRoles) TableName() string {
// 	return "UserRoles"
// }

func (m *UserRoles) GetID() interface{} {
	return m
}
