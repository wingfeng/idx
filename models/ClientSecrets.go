package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
)

// ClientSecrets [...]
type ClientSecrets struct {
	ID           int       `gorm:"primary_key;AUTO_INCREMENT;column:Id;not null"`
	Description  string    `gorm:"column:Description;type:varchar(2000)"`
	Value        string    `gorm:"column:Value;type:varchar(256);not null"`
	Expiration   time.Time `gorm:"column:Expiration;type:datetime(6)"`
	Type         string    `gorm:"column:Type;type:varchar(250);not null"`
	ClientID     int       `gorm:"index:IX_ClientSecrets_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Client    `gorm:"foreignkey:ClientId"`
	utils.Record `gorm:"embedded"`
}
