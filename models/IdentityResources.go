package models

// IdentityResources [...]
type IdentityResources struct {
	Id                      int                  `gorm:"primary_key;auto_Increment;not null"`
	Enabled                 bool                 `gorm:"not null"`
	Name                    string               `gorm:"unique;type:varchar(200);not null"`
	DisplayName             string               `gorm:"type:varchar(200)"`
	Description             string               `gorm:"type:varchar(1000)"`
	Required                bool                 `gorm:"not null"`
	Emphasize               bool                 `gorm:"not null"`
	ShowInDiscoveryDocument bool                 `gorm:"not null"`
	Properties              []IdentityProperties `gorm:"foreignkey:IdentityResourceId"`
	Claims                  []IdentityClaims     `gorm:"foreignkey:IdentityResourceId"`
	Record                  `gorm:"embedded"`
}
