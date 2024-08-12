package models

// ClientCorsOrigins [...]
type ClientCorsOrigins struct {
	Id       int    `gorm:"primary_key;auto_Increment;not null"`
	Origin   string `gorm:"type:varchar(150);not null"`
	ClientId int    `gorm:"index:IX_ClientCorsOrigins_ClientId;type:int;not null"`

	Record `gorm:"embedded"`
}
