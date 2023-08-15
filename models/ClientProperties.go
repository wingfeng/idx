package models

import "github.com/wingfeng/idx/utils"

// ClientProperties [...]
type ClientProperties struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:id;not null"`
	Key          string `gorm:"column:key;type:varchar(250);not null"`
	Value        string `gorm:"column:value;type:varchar(2000);not null"`
	ClientID     int    `gorm:"index:IX_ClientProperties_ClientId;column:clientid;type:int;not null"`
	Clients      Client `gorm:"foreignKey:clientid;references:id"`
	utils.Record `gorm:"embedded"`
}
