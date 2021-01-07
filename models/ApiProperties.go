package models

import "github.com/wingfeng/idx/utils"

// APIProperties [...]
type APIProperties struct {
	ID            int          `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Key           string       `gorm:"column:Key;type:varchar(250);not null"`
	Value         string       `gorm:"column:Value;type:varchar(2000);not null"`
	APIResourceID int          `gorm:"index:IX_ApiProperties_ApiResourceId;column:ApiResourceId;type:int(11);not null"`
	APIResources  APIResources `gorm:"association_foreignkey:ApiResourceId;foreignkey:Id"`
	utils.Record  `gorm:"embedded"`
}
