package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// APIResources [...]
type APIResources struct {
	ID           int       `gorm:"primary_key;autoIncrement;column:id;not null"`
	Enabled      bool      `gorm:"column:enabled;not null"`
	Name         string    `gorm:"unique;column:name;type:varchar(200);not null"`
	DisplayName  string    `gorm:"column:displayname;type:varchar(200)"`
	Description  string    `gorm:"column:description;type:varchar(1000)"`
	Created      time.Time `gorm:"column:created;not null"`
	Updated      null.Time `gorm:"column:updated;"`
	LastAccessed null.Time `gorm:"column:lastaccessed;"`
	NonEditable  bool      `gorm:"column:noneditable;not null"`
	utils.Record `gorm:"embedded"`
}
