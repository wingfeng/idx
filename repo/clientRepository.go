package repo

import (
	"github.com/wingfeng/idx-oauth2/model"
	"github.com/wingfeng/idx/models"

	"gorm.io/gorm"
)

type DBClientRepository struct {
	DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) *DBClientRepository {
	return &DBClientRepository{DB: db}
}
func (r *DBClientRepository) GetClientByClientID(id string) (model.IClient, error) {
	result := &models.Client{}
	tx := r.DB.Table(result.TableName()).Where("client_id = ?", id).Preload("Secrets").Preload("Origins").Preload("LogoutUris").First(result)
	return result, tx.Error
}
