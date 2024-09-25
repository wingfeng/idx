package models

// UserTokens [...]
type UserTokens struct {
	LoginProvider string `gorm:"primary_key;type:varchar(255);not null"`
	Name          string `gorm:"primary_key;type:varchar(255);not null"`
	Value         string `gorm:"type:text"`

	Record `gorm:"embedded"`
}
