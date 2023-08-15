package models

import "github.com/wingfeng/idx/utils"

// ClientClaims [...]
type ClientClaims struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:id;not null"`
	Type         string `gorm:"column:type;type:varchar(250);not null"`
	Value        string `gorm:"column:value;type:varchar(250);not null"`
	ClientID     int    `gorm:"index:IX_ClientClaims_ClientId;column:clientid;type:int;not null"`
	Clients      Client `gorm:"foreignKey:clientid;references:id"`
	utils.Record `gorm:"embedded"`
}
