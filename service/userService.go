package service

import (
	"github.com/wingfeng/idx-oauth2/service"
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
