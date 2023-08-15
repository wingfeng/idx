package models

import "github.com/wingfeng/idx/utils"

// APIScopeClaims [...]
type APIScopeClaims struct {
	ID           int       `gorm:"primary_key;auto_Increment;column:id;not null"`
	Type         string    `gorm:"column:type;type:varchar(200);not null"`
	APIScopeID   int       `gorm:"index:IX_ApiScopeClaims_ApiScopeId;column:apiscopeid;type:int;not null"`
	APIScopes    APIScopes `gorm:"foreignkey:apiscopeid"`
	utils.Record `gorm:"embedded"`
}
