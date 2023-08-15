package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// IDentityResources [...]
type IDentityResources struct {
	ID                      int       `gorm:"primary_key;auto_Increment;column:id;not null"`
	Enabled                 bool      `gorm:"column:enabled;not null"`
	Name                    string    `gorm:"unique;column:name;type:varchar(200);not null"`
	DisplayName             string    `gorm:"column:displayname;type:varchar(200)"`
	Description             string    `gorm:"column:description;type:varchar(1000)"`
	Required                bool      `gorm:"column:required;not null"`
	Emphasize               bool      `gorm:"column:emphasize;not null"`
	ShowInDiscoveryDocument bool      `gorm:"column:showindiscoverydocument;not null"`
	Created                 time.Time `gorm:"column:created;not null"`
	Updated                 null.Time `gorm:"column:updated;"`
	NonEditable             bool      `gorm:"column:noneditable;not null"`
	utils.Record            `gorm:"embedded"`
}
