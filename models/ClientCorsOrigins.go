package models

import "github.com/wingfeng/idx/utils"

// ClientCorsOrigins [...]
type ClientCorsOrigins struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:id;not null"`
	Origin       string `gorm:"column:origin;type:varchar(150);not null"`
	ClientID     int    `gorm:"index:IX_ClientCorsOrigins_ClientId;column:clientid;type:int;not null"`
	Clients      Client `gorm:"foreignKey:clientid;references:id"`
	utils.Record `gorm:"embedded"`
}
