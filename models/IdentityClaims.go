package models

import "github.com/wingfeng/idx/utils"

// IDentityClaims [...]
type IDentityClaims struct {
	ID                 int               `gorm:"primary_key;auto_Increment;column:id;not null"`
	Type               string            `gorm:"column:type;type:varchar(200);not null"`
	IDentityResourceID int               `gorm:"index:IX_IdentityClaims_IdentityResourceId;column:identityresourceid;type:int;not null"`
	IDentityResources  IDentityResources `gorm:"foreignkey:identityresourceid"`
	utils.Record       `gorm:"embedded"`
}
