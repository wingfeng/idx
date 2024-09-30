package dto

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/copier"
	"github.com/wingfeng/idx/models"
	"gorm.io/datatypes"
)

type UserDto struct {
	Id                   snowflake.ID   `json:"id"`
	OUId                 snowflake.ID   `json:"ouid"`
	OU                   string         `json:"ou" `
	Account              string         `json:"preferred_username" `
	DisplayName          string         `json:"displayname" `
	Email                string         `json:"email" `
	EmailConfirmed       bool           `json:"emailconfirmed"`
	PhoneNumber          string         `json:"phonenumber"`
	PhoneNumberConfirmed bool           `json:"phonenumberconfirmed"`
	PasswordHash         string         `json:"-"`
	IsTemporaryPassword  bool           `json:"istemporarypassword"`
	Lockoutenabled       bool           `json:"lockoutenabled"`
	Lockoutend           time.Time      `json:"lockoutend"`
	Roles                []string       `json:"roles" copier:"-" gorm:"-"`
	Claims               datatypes.JSON `json:"claims"`
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
	return r.Id.String()
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
