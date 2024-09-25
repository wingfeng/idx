package repo

import (
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/wingfeng/idx-oauth2/model"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/models/dto"
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
	tx := repo.DB.Model(models.User{}).Where("id = ?", p).First(&user)

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

	repo.DB.SetupJoinTable(&models.User{}, "Roles", &models.UserRoles{})
	tx := repo.DB.Where("normalized_account = ?", strings.ToUpper(username)).First(&user)
	if tx.Error == nil {
		repo.DB.Model(&roles).Joins("join user_roles on user_roles.role_id=roles.id").Where("user_roles.user_id=?", user.Id).Find(&roles)
		user.Roles = roles
	}

	dto := dto.NewUserDto(&user)
	return dto, tx.Error
}
