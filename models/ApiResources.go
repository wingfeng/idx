package models

import "gorm.io/datatypes"

// APIResources [...]
type APIResources struct {
	Enabled     bool   `gorm:"not null"`
	Name        string `gorm:"unique;type:varchar(200);not null"`
	DisplayName string `gorm:"type:varchar(200)"`
	Description string `gorm:"type:varchar(1000)"`
	Claims      datatypes.JSON
	Properties  datatypes.JSON
	IntRecord   `gorm:"embedded"`
}

func (APIResources) TableName() string {
	return "api_resources"
}
