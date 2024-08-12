package models

// IdentityClaims [...]
type IdentityClaims struct {
	Id                 int    `gorm:"primary_key;auto_Increment;not null"`
	Type               string `gorm:"type:varchar(200);not null"`
	IdentityResourceId int    `gorm:"index:IX_IdentityClaims_IdentityResourceId;column:identityresourceid;type:int;not null"`
	Record             `gorm:"embedded"`
}
