package models

import "github.com/wingfeng/idx/utils"

// APIProperties [...]
type APIProperties struct {
	ID            int          `gorm:"primary_key;autoIncrement;column:id;not null"`
	Key           string       `gorm:"column:key;type:varchar(250);not null"`
	Value         string       `gorm:"column:value;type:varchar(2000);not null"`
	APIResourceID int          `gorm:"index:IX_ApiProperties_ApiResourceId;column:apiresourceid;type:int;not null"`
	APIResources  APIResources `gorm:"foreignkey:apiresourceid"`
	utils.Record  `gorm:"embedded"`
}
