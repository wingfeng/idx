package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// APIResources [...]
type APIResources struct {
	ID           int       `gorm:"primary_key;autoIncrement;column:Id;not null"`
	Enabled      bool      `gorm:"column:Enabled;type:tinyint(1);not null"`
	Name         string    `gorm:"unique;column:Name;type:varchar(200);not null"`
	DisplayName  string    `gorm:"column:DisplayName;type:varchar(200)"`
	Description  string    `gorm:"column:Description;type:varchar(1000)"`
	Created      time.Time `gorm:"column:Created;type:datetime(6);not null"`
	Updated      null.Time `gorm:"column:Updated;type:datetime(6)"`
	LastAccessed null.Time `gorm:"column:LastAccessed;type:datetime(6)"`
	NonEditable  bool      `gorm:"column:NonEditable;type:tinyint(1);not null"`
	utils.Record `gorm:"embedded"`
}
