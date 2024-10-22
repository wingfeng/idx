package models

import "github.com/bwmarrin/snowflake"

// UserRoles [...]
type UserRoles struct {
	UserId snowflake.ID `gorm:"primary_key;not null"`
	RoleId snowflake.ID `gorm:"primary_key;not null"`

	//	Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *UserRoles) TableName() string {
// 	return "UserRoles"
// }

// func (m *UserRoles) GetID() interface{} {
// 	return m
// }
