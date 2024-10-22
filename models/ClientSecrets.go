package models

import (
	"time"
)

// ClientSecrets [...]
type ClientSecrets struct {
	Name       string    `gorm:"type:varchar(100)"`
	Value      string    `gorm:"type:varchar(256);not null"`
	Expiration time.Time `gorm:"column:expiration;"`

	ClientId  int64 `gorm:"type:varchar(256);not null"`
	IntRecord `gorm:"embedded"`
}
