package models

// ClientProperties [...]
type ClientProperties struct {
	Id       int    `gorm:"primary_key;auto_Increment;not null"`
	Key      string `gorm:"type:varchar(250);not null"`
	Value    string `gorm:"type:varchar(2000);not null"`
	ClientId int    `gorm:"index:IX_ClientProperties_ClientId;type:int;not null"`

	Record `gorm:"embedded"`
}
