package models

import (
	"strings"

	"github.com/bwmarrin/snowflake"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Role [...]
type Role struct {
	Name           string `gorm:"type:varchar(256)"`
	NormalizedName string `gorm:"unique;type:varchar(256)"`

	Users           []User `gorm:"many2many:user_roles;"`
	Claims          datatypes.JSON
	SnowflakeRecord `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *Role) TableName() string {
// 	return "Roles"
// }

// SetID 获取当前记录的ID
func (r *Role) SetID(v interface{}) {
	r.Id = v.(snowflake.ID)
}

func (r *Role) GetID() interface{} {
	return r.Id
}
func (r *Role) BeforeCreate(tx *gorm.DB) error {

	r.NormalizedName = strings.ToUpper(r.Name)

	return nil
}
func (r *Role) BeforeUpdate(tx *gorm.DB) error {

	r.NormalizedName = strings.ToUpper(r.Name)

	return nil
}
