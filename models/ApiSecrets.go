package models

import (
	"gopkg.in/guregu/null.v4"
)

// APISecrets [...]
type APISecrets struct {
	Description   string `gorm:"type:varchar(1000)"`
	Value         string `gorm:"type:text;not null"`
	Type          string `gorm:"type:varchar(250);not null"`
	APIResourceId int    `gorm:"index:IX_ApiSecrets_ApiResourceId;not null"`
	Expiration    null.Time
	IntRecord     `gorm:"embedded"`
}
