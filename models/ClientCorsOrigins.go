package models

import "github.com/wingfeng/idx/utils"

// ClientCorsOrigins [...]
type ClientCorsOrigins struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:Id;not null"`
	Origin       string `gorm:"column:Origin;type:varchar(150);not null"`
	ClientID     int    `gorm:"index:IX_ClientCorsOrigins_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Client `gorm:"foreignkey:ClientId"`
	utils.Record `gorm:"embedded"`
}
