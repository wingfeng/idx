package models

// APIClaims [...]
type APIClaims struct {
	Id            int    `gorm:"primary_key;auto_Increment;not null"`
	Type          string `gorm:"type:varchar(200);not null"`
	APIResourceId int    `gorm:"index:IX_ApiClaims_ApiResourceId;type:int;not null"`
	Record        `gorm:"embedded"`
}
