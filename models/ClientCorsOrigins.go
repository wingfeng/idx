package models

// ClientCorsOrigins [...]
type ClientCorsOrigins struct {
	Origin   string `gorm:"type:varchar(150);not null"`
	ClientId int64  `gorm:"index:IX_ClientCorsOrigins_ClientId;type:bigint;not null"`

	IntRecord `gorm:"embedded"`
}
