package models

import "gorm.io/datatypes"

// Scopes [...]
type Scopes struct {
	Enabled     bool   `gorm:"not null"`
	Name        string `gorm:"unique;type:varchar(200);not null"`
	DisplayName string `gorm:"type:varchar(200)"`
	Description string `gorm:"type:varchar(1000)"`
	Claims      datatypes.JSON
	Properties  datatypes.JSON
	IntRecord   `gorm:"embedded"`
}
