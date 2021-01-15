package models

import "github.com/wingfeng/idx/utils"

// ClientProperties [...]
type ClientProperties struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:Id;not null"`
	Key          string `gorm:"column:Key;type:varchar(250);not null"`
	Value        string `gorm:"column:Value;type:varchar(2000);not null"`
	ClientID     int    `gorm:"index:IX_ClientProperties_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Client `gorm:"foreignkey:ClientId"`
	utils.Record `gorm:"embedded"`
}
