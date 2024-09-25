package models

import "gorm.io/datatypes"

// APIScopes [...]
type APIScopes struct {
	Name                    string `gorm:"unique;type:varchar(200);not null"`
	DisplayName             string `gorm:"type:varchar(200)"`
	Description             string `gorm:"type:varchar(1000)"`
	Required                bool   `gorm:"not null"`
	Emphasize               bool   `gorm:"not null"`
	ShowInDiscoveryDocument bool   `gorm:"not null"`
	APIResourceID           int    `gorm:"index:IX_ApiScopes_ApiResourceId;column:apiresourceid;type:int;not null"`
	IntRecord               `gorm:"embedded"`
	Claims                  datatypes.JSON
}
