package models

import "github.com/wingfeng/idx/utils"

// IDentityProperties [...]
type IDentityProperties struct {
	ID                 int               `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Key                string            `gorm:"column:Key;type:varchar(250);not null"`
	Value              string            `gorm:"column:Value;type:varchar(2000);not null"`
	IDentityResourceID int               `gorm:"index:IX_IdentityProperties_IdentityResourceId;column:IdentityResourceId;type:int(11);not null"`
	IDentityResources  IDentityResources `gorm:"association_foreignkey:IdentityResourceId;foreignkey:Id"`
	utils.Record       `gorm:"embedded"`
}
