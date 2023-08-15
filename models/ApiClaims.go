package models

import "github.com/wingfeng/idx/utils"

// APIClaims [...]
type APIClaims struct {
	ID            int          `gorm:"primary_key;auto_Increment;column:id;not null"`
	Type          string       `gorm:"column:type;type:varchar(200);not null"`
	APIResourceID int          `gorm:"index:IX_ApiClaims_ApiResourceId;column:apiresourceid;type:int;not null"`
	APIResources  APIResources `gorm:"foreignkey:apiresourceid"`
	utils.Record  `gorm:"embedded"`
}
