package repo

import (
	"github.com/wingfeng/idx/models"
	"gorm.io/gorm"
)

type DBScopeRepository struct {
	DB *gorm.DB
}

func NewScopeRepository(db *gorm.DB) *DBScopeRepository {
	return &DBScopeRepository{
		DB: db,
	}
}
func (r *DBScopeRepository) GetSupportScopes() ([]string, error) {
	var scopes []models.Scopes
	tx := r.DB.Model(&models.Scopes{}).Where("enabled = ?", true).Find(&scopes)
	result := make([]string, 0)
	for _, scope := range scopes {
		result = append(result, scope.Name)
	}
	return result, tx.Error
}
