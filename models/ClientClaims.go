package models

import "github.com/wingfeng/idx/utils"

// ClientClaims [...]
type ClientClaims struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:Id;not null"`
	Type         string `gorm:"column:Type;type:varchar(250);not null"`
	Value        string `gorm:"column:Value;type:varchar(250);not null"`
	ClientID     int    `gorm:"index:IX_ClientClaims_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Client `gorm:"foreignkey:ClientId"`
	utils.Record `gorm:"embedded"`
}
