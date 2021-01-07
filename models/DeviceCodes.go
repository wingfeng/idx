package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
)

// DeviceCodes [...]
type DeviceCodes struct {
	UserCode     string    `gorm:"primary_key;column:UserCode;type:varchar(200);not null"`
	DeviceCode   string    `gorm:"unique;column:DeviceCode;type:varchar(200);not null"`
	SubjectID    string    `gorm:"column:SubjectId;type:varchar(200)"`
	ClientID     string    `gorm:"column:ClientId;type:varchar(200);not null"`
	CreationTime time.Time `gorm:"column:CreationTime;type:datetime(6);not null"`
	Expiration   time.Time `gorm:"index:IX_DeviceCodes_Expiration;column:Expiration;type:datetime(6);not null"`
	Data         string    `gorm:"column:Data;type:longtext;not null"`
	utils.Record `gorm:"embedded"`
}
