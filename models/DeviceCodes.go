package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
)

// DeviceCodes [...]
type DeviceCodes struct {
	UserCode     string    `gorm:"primary_key;column:usercode;type:varchar(200);not null"`
	DeviceCode   string    `gorm:"unique;column:devicecode;type:varchar(200);not null"`
	SubjectID    string    `gorm:"column:subjectid;type:varchar(200)"`
	ClientID     string    `gorm:"column:clientid;type:varchar(200);not null"`
	CreationTime time.Time `gorm:"column:creationtime;not null"`
	Expiration   time.Time `gorm:"index:IX_DeviceCodes_Expiration;column:expiration;not null"`
	Data         string    `gorm:"column:data;type:text;not null"`
	utils.Record `gorm:"embedded"`
}
