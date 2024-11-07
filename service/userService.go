package service

import (
	"github.com/wingfeng/idx-oauth2/model"
	"github.com/wingfeng/idx-oauth2/service"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/repo"
)

type UserService struct {
	repo *repo.DBUserRepository
	service.DefaultUserService
}

func NewUserService(repo *repo.DBUserRepository) *UserService {
	s := &UserService{
		repo: repo,
	}
	s.UserRepository = repo
	return s
}

func (s *UserService) ChangePassword(username, oldPassword, newPassword string) error {
	return s.repo.ChangePassword(username, oldPassword, newPassword)
}
func (s *UserService) SearchUsers(filters *models.Expression, skip, take int) ([]model.IUser, error) {
	re, err := s.repo.SearchUsers(filters, skip, take)
	result := make([]model.IUser, len(re))
	for i, v := range re {
		result[i] = &v
	}
	return result, err
}
