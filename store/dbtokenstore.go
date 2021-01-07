package store

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	"gorm.io/gorm"
)

//DBTokenStore 存储于数据库的Token
type DBTokenStore struct {
	DB *gorm.DB
}

//Create and store the new token information
func (dts *DBTokenStore) Create(ctx context.Context, info oauth2.TokenInfo) error {
	return nil
}

//RemoveByCode delete the authorization code
func (dts *DBTokenStore) RemoveByCode(ctx context.Context, code string) error {
	return nil
}

//RemoveByAccess use the access token to delete the token information
func (dts *DBTokenStore) RemoveByAccess(ctx context.Context, access string) error {
	return nil
}

//RemoveByRefresh use the refresh token to delete the token information
func (dts *DBTokenStore) RemoveByRefresh(ctx context.Context, refresh string) error {
	return nil
}

//GetByCode use the authorization code for token information data
func (dts *DBTokenStore) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	return nil, nil
}

//GetByAccess use the access token for token information data
func (dts *DBTokenStore) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	return nil, nil
}

//GetByRefresh use the refresh token for token information data
func (dts *DBTokenStore) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	return nil, nil
}
