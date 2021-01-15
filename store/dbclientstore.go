package store

import (
	"context"

	idxmodels "github.com/wingfeng/idx/models"
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
func (cs *ClientStore) GetByID(ctx context.Context, id string) (*idxmodels.Client, error) {
	client := &idxmodels.Client{}
	err := cs.DB.Where("ClientId=?", id).First(client).Error
	if err != nil {
		return nil, err
	}
	return client, err
}

func (cs *ClientStore) GetClientRedirectUris(id int) ([]string, error) {
	uris := new([]idxmodels.ClientRedirectURIs)
	err := cs.DB.Where("ClientId=?", id).Find(uris).Error
	var sURIs []string
	for _, s := range *uris {
		sURIs = append(sURIs, s.RedirectURI)
	}
	return sURIs, err
}
func (cs *ClientStore) ValidateSecret(secret string) error {
	//验证Secret
	return nil
}

func (cs *ClientStore) GetClientScopes(clientID string) []string {

	clientScopes := []idxmodels.ClientScopes{}
	cs.DB.Table("client_scopes  as cs").Joins("join clients as c on cs.ClientId=c.Id ").Where("c.ClientId=?", clientID).Find(&clientScopes)
	var result []string
	for _, cs := range clientScopes {
		result = append(result, cs.Scope)
	}
	return result
}
