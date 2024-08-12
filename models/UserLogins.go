package models

// UserLogins [...]
type UserLogins struct {
	LoginProvider       string `gorm:"primary_key;type:varchar(255);not null"`
	ProviderKey         string `gorm:"primary_key;type:varchar(255);not null"`
	ProviderDisplayName string `gorm:"type:text"`
	UserId              string `gorm:"index:IX_UserLogins_UserId;type:varchar(255);not null"`
	Record              `gorm:"embedded"`
}
