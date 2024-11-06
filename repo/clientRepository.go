package repo

import (
	"context"
	"strings"

	"github.com/wingfeng/idx-oauth2/model"
	"github.com/wingfeng/idx/cache"
	"github.com/wingfeng/idx/consts"
	"github.com/wingfeng/idx/models"

	"gorm.io/gorm"
)

type DBClientRepository struct {
	DB *gorm.DB
}

func NewClientRepository() *DBClientRepository {
	db := DefaultDB()
	return &DBClientRepository{DB: db}
}
func (r *DBClientRepository) GetClientByClientID(id string) (model.IClient, error) {
	funcLoad := func(ctx context.Context) (interface{}, error) {
		result := &models.Client{}
		tx := r.DB.Table(result.TableName()).Where("client_id = ? and enabled=true", id).Preload("Secrets").Preload("Origins").Preload("LogoutUris").First(result)
		return result, tx.Error
	}
	key := consts.CONST_CLIENTKEY + strings.ToLower(id)
	re, err := cache.Service.GetOrSetFuncLock(context.Background(), key, funcLoad, consts.CONST_TTL)
	if err == nil {
		result := &models.Client{}
		re.Struct(result)
		return result, err
	}

	return nil, err
}
