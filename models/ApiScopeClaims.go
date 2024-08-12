package models

// APIScopeClaims [...]
type APIScopeClaims struct {
	Id         int    `gorm:"primary_key;auto_Increment;not null"`
	Type       string `gorm:"type:varchar(200);not null"`
	APIScopeId int    `gorm:"index:IX_ApiScopeClaims_ApiScopeId;column:apiscopeid;type:int;not null"`

	Record `gorm:"embedded"`
}
