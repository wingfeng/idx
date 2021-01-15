package models

import "github.com/wingfeng/idx/utils"

// APIProperties [...]
type APIProperties struct {
	ID            int          `gorm:"primary_key;autoIncrement;column:Id;not null"`
	Key           string       `gorm:"column:Key;type:varchar(250);not null"`
	Value         string       `gorm:"column:Value;type:varchar(2000);not null"`
	APIResourceID int          `gorm:"index:IX_ApiProperties_ApiResourceId;column:ApiResourceId;type:int(11);not null"`
	APIResources  APIResources `gorm:"foreignkey:ApiResourceId"`
	utils.Record  `gorm:"embedded"`
}
