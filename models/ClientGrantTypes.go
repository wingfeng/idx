package models

import "github.com/wingfeng/idx/utils"

// ClientGrantTypes [...]
type ClientGrantTypes struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:Id;not null"`
	GrantType    string `gorm:"column:GrantType;type:varchar(250);not null"`
	ClientID     int    `gorm:"index:IX_ClientGrantTypes_ClientId;column:ClientId;not null"`
	Clients      Client `gorm:"foreignkey:ClientId;"`
	utils.Record `gorm:"embedded"`
}
