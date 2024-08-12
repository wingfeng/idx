package repo

import (
	"log/slog"

	"github.com/wingfeng/idx-oauth2/model"
	"gorm.io/gorm"
)

type DBAuthorizationRepository struct {
	DB *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) *DBAuthorizationRepository {
	return &DBAuthorizationRepository{
		DB: db,
	}
}

func (ar *DBAuthorizationRepository) Save(authorization *model.Authorization) {
	ar.DB.Save(authorization)
}
func (ar *DBAuthorizationRepository) Remove(authorization *model.Authorization) {
	ar.DB.Delete(authorization)
}
func (ar *DBAuthorizationRepository) GetAuthorizationByAccessToken(token string) *model.Authorization {

	result := &model.Authorization{}
	tx := ar.DB.Where("access_token = ?", token).First(result)
	if tx.Error != nil {
		slog.Error("GetAuthorizationByAccessToken Error", "error", tx.Error)
		return nil
	}
	return result
}
func (ar *DBAuthorizationRepository) GetAuthorizationByCode(code string) *model.Authorization {

	result := &model.Authorization{}
	tx := ar.DB.Where("code = ?", code).First(result)
	if tx.Error != nil {
		slog.Error("GetAuthorizationByCode Error", "error", tx.Error)
		return nil
	}
	return result
}
func (ar *DBAuthorizationRepository) GetAuthorizationByRefreshToken(token string) *model.Authorization {
	result := &model.Authorization{}
	tx := ar.DB.Where("refresh_token = ?", token).First(result)
	if tx.Error != nil {
		slog.Error("GetAuthorizationByRefreshToken Error", "error", tx.Error)
		return nil
	}
	return result
}
func (ar *DBAuthorizationRepository) GetAuthorizationByDeviceCode(device_code string) *model.Authorization {
	result := &model.Authorization{}
	tx := ar.DB.Where("device_code = ?", device_code).First(result)
	if tx.Error != nil {
		slog.Error("GetAuthorizationByDeviceCode Error", "error", tx.Error)
		return nil
	}
	return result
}
func (ar *DBAuthorizationRepository) GetAuthorizationByUserCode(user_code string) *model.Authorization {
	result := &model.Authorization{}
	tx := ar.DB.Where("user_code = ?", user_code).First(result)
	if tx.Error != nil {
		slog.Error("GetAuthorizationByUserCode Error", "error", tx.Error)
		return nil
	}
	return result
}
