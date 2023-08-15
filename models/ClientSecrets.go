package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
)

// ClientSecrets [...]
type ClientSecrets struct {
	ID           int       `gorm:"primary_key;AUTO_INCREMENT;column:id;not null"`
	Description  string    `gorm:"column:description;type:varchar(2000)"`
	Value        string    `gorm:"column:value;type:varchar(256);not null"`
	Expiration   time.Time `gorm:"column:expiration;"`
	Type         string    `gorm:"column:type;type:varchar(250);not null"`
	ClientID     int       `gorm:"index:IX_ClientSecrets_ClientId;column:clientid;type:int;not null"`
	Clients      Client    `gorm:"foreignKey:clientid;references:id"`
	utils.Record `gorm:"embedded"`
}
