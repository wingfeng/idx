package store

import (
	"strings"

	"github.com/wingfeng/idx/models"
	"gorm.io/gorm"
)

type DbUserStore struct {
	db *gorm.DB
}

func NewDbUserStore(db *gorm.DB) *DbUserStore {
	ds := &DbUserStore{
		db: db,
	}
	return ds
}
func (ds *DbUserStore) GetUserByAccount(account string) (*models.User, error) {
	user := &models.User{}
	account = strings.ToUpper(account)
	err := ds.db.Where("NormalizedAccount=?", account).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
func (ds *DbUserStore) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}

	err := ds.db.Where("Id=?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
func (ds *DbUserStore) GetUserPasswordHash(account string) (string, error) {
	user := &models.User{}

	err := ds.db.Select("PasswordHash").Where("Account=?", account).First(user).Error
	if err != nil {
		return "", err
	}
	return user.PasswordHash, err

}
