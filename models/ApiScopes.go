package models

// APIScopes [...]
type APIScopes struct {
	Id                      int    `gorm:"primary_key;auto_Increment;not null"`
	Name                    string `gorm:"unique;type:varchar(200);not null"`
	DisplayName             string `gorm:"type:varchar(200)"`
	Description             string `gorm:"type:varchar(1000)"`
	Required                bool   `gorm:"not null"`
	Emphasize               bool   `gorm:"not null"`
	ShowInDiscoveryDocument bool   `gorm:"not null"`
	APIResourceID           int    `gorm:"index:IX_ApiScopes_ApiResourceId;column:apiresourceid;type:int;not null"`
	Record                  `gorm:"embedded"`
}
