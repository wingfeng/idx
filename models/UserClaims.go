package models

// UserClaims [...]
type UserClaims struct {
	Id         int    `gorm:"primary_key;auto_Increment;not null"`
	UserId     string `gorm:"index:IX_UserClaims_UserId;type:varchar(255);not null"`
	ClaimType  string `gorm:"type:text"`
	ClaimValue string `gorm:"type:text"`
	Record     `gorm:"embedded"`
}
