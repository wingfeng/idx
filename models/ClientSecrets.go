package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
)

// ClientSecrets [...]
type ClientSecrets struct {
	ID           int       `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Description  string    `gorm:"column:Description;type:varchar(2000)"`
	Value        string    `gorm:"column:Value;type:longtext;not null"`
	Expiration   time.Time `gorm:"column:Expiration;type:datetime(6)"`
	Type         string    `gorm:"column:Type;type:varchar(250);not null"`
	ClientID     int       `gorm:"index:IX_ClientSecrets_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Client    `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	Created      time.Time `gorm:"column:Created"`
	utils.Record `gorm:"embedded"`
}
