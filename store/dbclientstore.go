package store

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cihub/seelog"
	"github.com/wingfeng/idx/cache"
	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/oauth2"
	"github.com/wingfeng/idx/utils"
	"gorm.io/gorm"
)

// ClientStore  用于存储Client信息的
type ClientStore struct {
	DB    *gorm.DB
	Cache cache.ICacheProvider
	mutex sync.RWMutex
}

func NewClientStore(db *gorm.DB) *ClientStore {
	cs := &ClientStore{
		DB: db,
	}
	return cs
}

// GetByID 通过ID获取Client信息
func (cs *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	var client *idxmodels.Client
	var err error
	if cs.Cache != nil {
		cs.mutex.RLock()
		obj, exist := cs.Cache.Get(id)
		cs.mutex.RUnlock()
		if exist && obj != nil {
			seelog.Tracef("Load from cache")

			client = obj.(*idxmodels.Client)
		} else {
			cs.mutex.Lock()
			client, err = cs.getFromDB(id)
			cs.Cache.Set(id, client, 2*time.Minute)

			cs.mutex.Unlock()
		}
	} else {
		client, err = cs.getFromDB(id)
	}
	return client, err
}
func (cs *ClientStore) getFromDB(id string) (*idxmodels.Client, error) {
	client := &idxmodels.Client{}
	seelog.Tracef("Load client from db")
	client.ClientCode = id
	client.Enabled = true
	err := cs.DB.Where(client).First(client).Error
	if err != nil {
		seelog.Errorf("DB Error :%v", err)
		return nil, err
	}
	return client, nil
}
func (cs *ClientStore) ValidateSecret(clientId, secret string) error {
	key := utils.HashString(secret)
	var count int64
	err := cs.DB.Table("client_secrets as cs ").Select("Count(1)").Joins("join clients as c on cs.ClientId=c.Id").
		Where("c.clientcode=? and cs.value=? and cs.expiration>Now()", clientId, key).Count(&count).Error
	if err == nil && count < 1 {
		err = fmt.Errorf("Client %s Secret not validate", clientId)
	}
	//验证Secret
	return err
}

// func (cs *ClientStore) GetClientScopes(clientID string) []string {

// 	clientScopes := []idxmodels.ClientScopes{}
// 	cs.DB.Table("client_scopes  as cs").Joins("join clients as c on cs.ClientId=c.Id ").Where("c.ClientId=?", clientID).Find(&clientScopes)
// 	var result []string
// 	for _, cs := range clientScopes {
// 		result = append(result, cs.Scope)
// 	}
// 	return result
// }
