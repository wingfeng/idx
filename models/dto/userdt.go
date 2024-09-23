package dto

import (
	"github.com/jinzhu/copier"
	"github.com/wingfeng/idx/models"
)

type UserDto struct {
	Id           string   `json:"id"`
	OUId         string   `json:"ouid"`
	OU           string   `json:"ou" `
	Account      string   `json:"preferred_username" `
	DisplayName  string   `json:"displayname" `
	Email        string   `json:"email" `
	PhoneNumber  string   `json:"phonenumber"`
	PasswordHash string   `json:"-"`
	Roles        []string `json:"roles" copier:"-"`
	UserClaims   []string `json:"claims" copier:"-"`
}

func NewUserDto(user *models.User) *UserDto {
	dto := &UserDto{}
	copier.Copy(dto, user)
	rolesMap := map[string]bool{}
	roles := []string{}
	for _, role := range user.Roles {
		if !rolesMap[role.Name] {
			rolesMap[role.Name] = true
			roles = append(roles, role.Name)
		}
	}
	dto.Roles = roles
	return dto
}

func (r *UserDto) GetId() string {
	return r.Id
}

func (r *UserDto) GetUserName() string {
	return r.Account
}
func (r *UserDto) GetEmail() string {
	return r.Email
}
func (r *UserDto) GetPasswordHash() string {
	return r.PasswordHash
}
