package models

// IdentityClaims [...]
type IdentityClaims struct {
	Type               string `gorm:"type:varchar(200);not null"`
	IdentityResourceId int    `gorm:"index:IX_IdentityClaims_IdentityResourceId;column:identityresourceid;type:int;not null"`
	IntRecord          `gorm:"embedded"`
}
