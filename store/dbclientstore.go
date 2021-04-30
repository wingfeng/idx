package store

import (
	"context"
	"fmt"

	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/oauth2"
	"github.com/wingfeng/idx/utils"
	"gorm.io/gorm"
)

//ClientStore  用于存储Client信息的
type ClientStore struct {
	DB *gorm.DB
}

func NewClientStore(db *gorm.DB) *ClientStore {
	cs := &ClientStore{
		DB: db,
	}
	return cs
}

//GetByID 通过ID获取Client信息
func (cs *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	client := &idxmodels.Client{}
	err := cs.DB.Where("ClientId=?", id).First(client).Error
	if err != nil {
		return nil, err
	}
	return client, err
}

func (cs *ClientStore) ValidateSecret(clientId, secret string) error {
	key := utils.HashString(secret)
	var count int64
	err := cs.DB.Table("client_secrets as cs ").Select("Count(1)").Joins("join clients as c on cs.ClientId=c.Id").
		Where("c.ClientId=? and cs.Value=? and cs.Expiration>Now()", clientId, key).Count(&count).Error
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
