package models

// ClientClaims [...]
type ClientClaims struct {
	Id       int    `gorm:"primary_key;auto_Increment;not null"`
	Type     string `gorm:"type:varchar(250);not null"`
	Value    string `gorm:"type:varchar(250);not null"`
	ClientId int    `gorm:"index:IX_ClientClaims_ClientId;type:int;not null"`
	Record   `gorm:"embedded"`
}
