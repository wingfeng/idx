package models

// APIResources [...]
type APIResources struct {
	Id          int    `gorm:"primary_key;autoIncrement;not null"`
	Enabled     bool   `gorm:"not null"`
	Name        string `gorm:"unique;type:varchar(200);not null"`
	DisplayName string `gorm:"type:varchar(200)"`
	Description string `gorm:"type:varchar(1000)"`

	Record `gorm:"embedded"`
}
