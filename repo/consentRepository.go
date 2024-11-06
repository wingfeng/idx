package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/wingfeng/idx/cache"
	"github.com/wingfeng/idx/consts"
	"github.com/wingfeng/idx/models"
	"gorm.io/gorm"
)

type DBConsentRepository struct {
	DB *gorm.DB
}

func NewConsentRepository() *DBConsentRepository {
	db := DefaultDB()
	return &DBConsentRepository{DB: db}
}

func (r *DBConsentRepository) GetConsents(clientId string, principal string) ([]string, error) {
	funcLoad := func(ctx context.Context) (interface{}, error) {
		result := &models.PersistedGrants{}
		tx := r.DB.Table(result.TableName()).Where("client_id = ? and principal = ?", clientId, principal).First(result)
		if tx.Error != nil {
			return nil, tx.Error
		}
		scope := strings.Split(result.Scope, " ")
		return scope, nil
	}
	key := fmt.Sprintf(consts.CONST_ConsentKey, clientId, principal)

	re, err := cache.Service.GetOrSetFuncLock(context.Background(), key, funcLoad, consts.CONST_TTL)
	if err == nil {
		return re.Strings(), nil
	}
	return nil, err
}
func (r *DBConsentRepository) SaveConsents(clientId string, principal string, scopes []string) error {
	if len(scopes) > 0 {
		scope := strings.Join(scopes, " ")
		result := &models.PersistedGrants{
			ClientId:  clientId,
			Principal: principal,
			Scope:     scope,
		}

		tx := r.DB.Table(result.TableName()).Where("client_id = ? and principal = ?", clientId, principal).Save(result)
		return tx.Error
	}
	return errors.New("scope is empty")
}
func (r *DBConsentRepository) RemoveConsents(clientId string, principal string) error {
	result := &models.PersistedGrants{
		ClientId:  clientId,
		Principal: principal,
	}

	err := r.DB.Table(result.TableName()).Delete(result).Error
	if err != nil {
		key := fmt.Sprintf(consts.CONST_ConsentKey, clientId, principal)
		cache.Service.Remove(context.Background(), key)
	}
	return err

}
