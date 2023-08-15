package models

import "github.com/wingfeng/idx/utils"

// IDentityProperties [...]
type IDentityProperties struct {
	ID                 int               `gorm:"primary_key;auto_Increment;column:id;not null"`
	Key                string            `gorm:"column:key;type:varchar(250);not null"`
	Value              string            `gorm:"column:value;type:varchar(2000);not null"`
	IDentityResourceID int               `gorm:"index:IX_IdentityProperties_IdentityResourceId;column:identityresourceid;type:int;not null"`
	IDentityResources  IDentityResources `gorm:"foreignkey:identityresourceid"`
	utils.Record       `gorm:"embedded"`
}
