package models

import (
	"strings"

	"github.com/wingfeng/idx/utils"
	"gorm.io/gorm"
)

// Role [...]
type Role struct {
	ID             string `json:"id" gorm:"primary_key;column:id;type:varchar(255);not null"`
	Name           string `gorm:"column:name;type:varchar(256)"`
	NormalizedName string `gorm:"unique;column:normalizedname;type:varchar(256)"`

	utils.Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *Role) TableName() string {
// 	return "Roles"
// }

// SetID 获取当前记录的ID
func (r *Role) SetID(v interface{}) {
	r.ID = v.(string)
}

func (r *Role) GetID() interface{} {
	return r.ID
}
func (r *Role) BeforeCreate(tx *gorm.DB) error {

	r.NormalizedName = strings.ToUpper(r.Name)

	return nil
}
func (r *Role) BeforeUpdate(tx *gorm.DB) error {

	r.NormalizedName = strings.ToUpper(r.Name)

	return nil
}
