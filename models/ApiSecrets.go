package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// APISecrets [...]
type APISecrets struct {
	ID            int          `gorm:"primary_key;auto_Increment;column:id;not null"`
	Description   string       `gorm:"column:description;type:varchar(1000)"`
	Value         string       `gorm:"column:value;type:text;not null"`
	Expiration    null.Time    `gorm:"column:expiration"`
	Type          string       `gorm:"column:type;type:varchar(250);not null"`
	Created       time.Time    `gorm:"column:created;not null"`
	APIResourceID int          `gorm:"index:IX_ApiSecrets_ApiResourceId;column:apiresourceid;not null"`
	APIResources  APIResources `gorm:"foreignkey:apiresourceid;"`
	utils.Record  `gorm:"embedded"`
}
