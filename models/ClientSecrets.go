package models

import (
	"time"
)

// ClientSecrets [...]
type ClientSecrets struct {
	Description string    `gorm:"type:varchar(2000)"`
	Value       string    `gorm:"type:varchar(256);not null"`
	Expiration  time.Time `gorm:"column:expiration;"`
	Type        string    `gorm:"type:varchar(250);not null"`
	ClientId    int64     `gorm:"type:varchar(256);not null"`
	IntRecord   `gorm:"embedded"`
}
