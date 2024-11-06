package repo

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/wingfeng/idx-oauth2/model"
	"github.com/wingfeng/idx-oauth2/utils"
	"github.com/wingfeng/idx/cache"
	_ "github.com/wingfeng/idx/cache" //make sure cache adapter can be loaded
	"github.com/wingfeng/idx/consts"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/models/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DBUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *DBUserRepository {
	db := DefaultDB()
	return &DBUserRepository{DB: db}
}
func (repo *DBUserRepository) GetUser(userId string) (model.IUser, error) {
	loadFunc := func(ctx context.Context) (interface{}, error) {
		var user models.User
		roles := make([]models.Role, 0)
		p, err := snowflake.ParseString(userId)
		if err != nil {
			return nil, err
		}
		tx := repo.DB.Model(&models.User{}).Preload("Roles").Where("id = ? and lockout_enabled=false", p).First(&user)

		if tx.Error == nil {
			repo.DB.Model(&roles).Joins("join user_roles on user_roles.role_id=roles.id").Where("user_roles.user_id=?", user.Id).Find(&roles)
			user.Roles = roles
		}

		return &user, tx.Error
	}

	cacheKey := consts.CONST_UserIDKey + userId
	r := &models.User{}

	val, err := cache.Service.GetOrSetFunc(context.Background(), cacheKey, loadFunc, consts.CONST_TTL)
	val.Struct(r)
	dto := dto.NewUserDto(r)
	return dto, err
}
func (repo *DBUserRepository) GetUserByName(username string) (model.IUser, error) {
	loadFunc := func(ctx context.Context) (interface{}, error) {
		var user models.User
		roles := make([]models.Role, 0)

		tx := repo.DB.Model(&models.User{}).Preload("Roles").Where("user_name = ? and lockout_enabled=false", strings.ToLower(username)).First(&user)
		if tx.Error == nil {
			repo.DB.Model(&roles).Joins("join user_roles on user_roles.role_id=roles.id").Where("user_roles.user_id=?", user.Id).Find(&roles)
			user.Roles = roles
		}
		slog.Info("load user from db")

		return user, tx.Error
	}
	cacheKey := fmt.Sprintf("%s%s", consts.CONST_USERNAMEKEY, strings.ToLower(username))

	user := &models.User{}

	val, err := cache.Service.GetOrSetFunc(context.Background(), cacheKey, loadFunc, consts.CONST_TTL)
	val.Struct(user)
	dto := dto.NewUserDto(user)
	return dto, err
}

func (repo *DBUserRepository) ChangePassword(username, oldPassword, newPassword string) error {
	//validate old password
	upassWord := struct {
		UserName     string
		PasswordHash string
	}{}
	tx := repo.DB.Model(&models.User{}).Where("user_name = ?", username).First(&upassWord)
	if tx.Error != nil {
		return tx.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(upassWord.PasswordHash), []byte(oldPassword))
	if err != nil {
		return err
	}
	newHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	u := map[string]interface{}{
		"password_hash":         newHash,
		"is_temporary_password": false,
	}
	err = repo.DB.Model(&models.User{}).Where("user_name = ?", username).Updates(u).Error
	if err == nil {
		//	go func() {
		ctx := context.Background()
		userKey := consts.CONST_USERNAMEKEY + username

		_, err := cache.Service.Remove(ctx, userKey)
		if err != nil {
			slog.Error("Remove cache error", "error", err)
		}
		key := consts.CONST_USERPWDHashKEY + strings.ToLower(username)
		_, err = cache.Service.Remove(ctx, key)
		if err != nil {
			slog.Error("Remove cache pwd hash error", "error", err)
		}
		//		}()
	}
	return err
	//update password

}

// Reset user password with a temporary password
func (repo *DBUserRepository) ResetPassword(username string) (string, error) {
	//generate password with random string
	newPassword := utils.GenerateRandomString(8)
	newHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return "", err
	}
	u := map[string]interface{}{
		"password_hash":         newHash,
		"is_temporary_password": true,
	}
	err = repo.DB.Model(&models.User{}).Where("user_name = ?", username).Updates(u).Error
	if err != nil {
		go func() {
			ctx := context.Background()
			key := consts.CONST_USERPWDHashKEY + strings.ToLower(username)
			cache.Service.Remove(ctx, key)
			key = consts.CONST_USERNAMEKEY + strings.ToLower(username)
			cache.Service.Remove(ctx, key)
		}()
	}
	return newPassword, err
}

func (repo *DBUserRepository) GetUserPasswordHash(username string) (string, error) {
	funcLoad := func(ctx context.Context) (interface{}, error) {
		s := ""
		tx := repo.DB.Model(&models.User{}).Select("password_hash").Where("user_name=?", username).First(&s)
		return s, tx.Error
	}
	key := consts.CONST_USERPWDHashKEY + strings.ToLower(username)
	r, err := cache.Service.GetOrSetFuncLock(context.Background(), key, funcLoad, consts.CONST_TTL)
	return r.String(), err

}
