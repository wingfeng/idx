package models

import "github.com/wingfeng/idx/utils"

// APIScopes [...]
type APIScopes struct {
	ID                      int          `gorm:"primary_key;auto_Increment;column:Id;not null"`
	Name                    string       `gorm:"unique;column:Name;type:varchar(200);not null"`
	DisplayName             string       `gorm:"column:DisplayName;type:varchar(200)"`
	Description             string       `gorm:"column:Description;type:varchar(1000)"`
	Required                bool         `gorm:"column:Required;type:tinyint(1);not null"`
	Emphasize               bool         `gorm:"column:Emphasize;type:tinyint(1);not null"`
	ShowInDiscoveryDocument bool         `gorm:"column:ShowInDiscoveryDocument;type:tinyint(1);not null"`
	APIResourceID           int          `gorm:"index:IX_ApiScopes_ApiResourceId;column:ApiResourceId;type:int(11);not null"`
	APIResources            APIResources `gorm:"foreignkey:ApiResourceId"`
	utils.Record            `gorm:"embedded"`
}
