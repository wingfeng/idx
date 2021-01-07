package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// APISecrets [...]
type APISecrets struct {
	ID            int          `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Description   string       `gorm:"column:Description;type:varchar(1000)"`
	Value         string       `gorm:"column:Value;type:longtext;not null"`
	Expiration    null.Time    `gorm:"column:Expiration;type:datetime(6)"`
	Type          string       `gorm:"column:Type;type:varchar(250);not null"`
	Created       time.Time    `gorm:"column:Created;type:datetime(6);not null"`
	APIResourceID int          `gorm:"index:IX_ApiSecrets_ApiResourceId;column:ApiResourceId;type:int(11);not null"`
	APIResources  APIResources `gorm:"association_foreignkey:ApiResourceId;foreignkey:Id"`
	utils.Record  `gorm:"embedded"`
}
