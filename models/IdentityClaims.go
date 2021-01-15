package models

import "github.com/wingfeng/idx/utils"

// IDentityClaims [...]
type IDentityClaims struct {
	ID                 int               `gorm:"primary_key;auto_Increment;column:Id;not null"`
	Type               string            `gorm:"column:Type;type:varchar(200);not null"`
	IDentityResourceID int               `gorm:"index:IX_IdentityClaims_IdentityResourceId;column:IdentityResourceId;type:int(11);not null"`
	IDentityResources  IDentityResources `gorm:"foreignkey:IdentityResourceId"`
	utils.Record       `gorm:"embedded"`
}
