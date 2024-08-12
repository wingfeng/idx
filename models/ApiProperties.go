package models

// APIProperties [...]
type APIProperties struct {
	Id            int    `gorm:"primary_key;autoIncrement;not null"`
	Key           string `gorm:"type:varchar(250);not null"`
	Value         string `gorm:"type:varchar(2000);not null"`
	APIResourceId int    `gorm:"index:IX_ApiProperties_ApiResourceId;type:int;not null"`
	Record        `gorm:"embedded"`
}
