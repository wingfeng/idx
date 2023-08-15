package models

import "github.com/wingfeng/idx/utils"

// APIScopes [...]
type APIScopes struct {
	ID                      int          `gorm:"primary_key;auto_Increment;column:id;not null"`
	Name                    string       `gorm:"unique;column:name;type:varchar(200);not null"`
	DisplayName             string       `gorm:"column:displayname;type:varchar(200)"`
	Description             string       `gorm:"column:description;type:varchar(1000)"`
	Required                bool         `gorm:"column:required;not null"`
	Emphasize               bool         `gorm:"column:emphasize;not null"`
	ShowInDiscoveryDocument bool         `gorm:"column:showindiscoverydocument;not null"`
	APIResourceID           int          `gorm:"index:IX_ApiScopes_ApiResourceId;column:apiresourceid;type:int;not null"`
	APIResources            APIResources `gorm:"foreignkey:apiresourceid"`
	utils.Record            `gorm:"embedded"`
}
