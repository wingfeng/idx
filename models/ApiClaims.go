package models

import "github.com/wingfeng/idx/utils"

// APIClaims [...]
type APIClaims struct {
	ID            int          `gorm:"primary_key;auto_increment;column:Id;type:int(11);not null"`
	Type          string       `gorm:"column:Type;type:varchar(200);not null"`
	APIResourceID int          `gorm:"index:IX_ApiClaims_ApiResourceId;column:ApiResourceId;type:int(11);not null"`
	APIResources  APIResources `gorm:"association_foreignkey:ApiResourceId;foreignkey:Id"`
	utils.Record  `gorm:"embedded"`
}
