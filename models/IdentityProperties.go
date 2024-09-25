package models

// IDentityProperties [...]
type IdentityProperties struct {
	Key                string `gorm:"type:varchar(250);not null"`
	Value              string `gorm:"type:varchar(2000);not null"`
	IdentityResourceId int    `gorm:"index:IX_IdentityProperties_IdentityResourceId;type:int;not null"`
	IntRecord          `gorm:"embedded"`
}
