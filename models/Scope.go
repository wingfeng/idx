package models

import "gorm.io/datatypes"

// oAuth2 Scopes
type Scopes struct {
	Enabled     bool   `gorm:"not null"`
	Name        string `gorm:"unique;type:varchar(200);not null"` //shold be lower case
	Description string `gorm:"type:varchar(1000)"`
	Properties  datatypes.JSON
	IntRecord   `gorm:"embedded"`
}
