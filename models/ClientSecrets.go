package models

import (
	"time"
)

// ClientSecrets [...]
type ClientSecrets struct {
	Id          int       `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Description string    `gorm:"type:varchar(2000)"`
	Value       string    `gorm:"type:varchar(256);not null"`
	Expiration  time.Time `gorm:"column:expiration;"`
	Type        string    `gorm:"type:varchar(250);not null"`
	ClientId    int       `gorm:"type:varchar(256);not null"`
	Record      `gorm:"embedded"`
}
