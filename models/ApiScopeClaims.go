package models

import "github.com/wingfeng/idx/utils"

// APIScopeClaims [...]
type APIScopeClaims struct {
	ID           int       `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Type         string    `gorm:"column:Type;type:varchar(200);not null"`
	APIScopeID   int       `gorm:"index:IX_ApiScopeClaims_ApiScopeId;column:ApiScopeId;type:int(11);not null"`
	APIScopes    APIScopes `gorm:"association_foreignkey:ApiScopeId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}
