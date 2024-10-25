package repo

import (
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/wingfeng/idx-oauth2/model"
	"github.com/wingfeng/idx-oauth2/utils"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/models/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DBUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *DBUserRepository {
	return &DBUserRepository{DB: db}
}
func (repo *DBUserRepository) GetUser(userId string) (model.IUser, error) {
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
	dto := dto.NewUserDto(&user)
	return dto, tx.Error
}
func (repo *DBUserRepository) GetUserByName(username string) (model.IUser, error) {
	var user models.User
	roles := make([]models.Role, 0)

	tx := repo.DB.Model(&models.User{}).Preload("Roles").Where("user_name = ? and lockout_enabled=false", strings.ToLower(username)).First(&user)
	if tx.Error == nil {
		repo.DB.Model(&roles).Joins("join user_roles on user_roles.role_id=roles.id").Where("user_roles.user_id=?", user.Id).Find(&roles)
		user.Roles = roles
	}

	dto := dto.NewUserDto(&user)
	return dto, tx.Error
}

func (repo *DBUserRepository) ChangePassword(username, oldPassword, newPassword string) error {
	//validate old password
	upassWord := struct {
		Account      string
		PasswordHash string
	}{}
	tx := repo.DB.Model(&models.User{}).Where("account = ?", username).First(&upassWord)
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
	return repo.DB.Model(&models.User{}).Where("account = ?", username).Updates(u).Error
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
	return newPassword, repo.DB.Model(&models.User{}).Where("account = ?", username).Updates(u).Error
}
