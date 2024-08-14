package models

import (
	"strings"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

// User 用户信息
type User struct {
	Id                   string    `json:"id" gorm:"primary_key;type:varchar(255);not null"`
	OUId                 string    `json:"ouid" gorm:"column:ou_id;type:varchar(36)"`
	OU                   string    `json:"ou" gorm:"type:varchar(256)"`
	Account              string    `json:"account" gorm:"type:varchar(256)"`
	DisplayName          string    `json:"displayname" gorm:"type:varchar(256)"`
	NormalizedAccount    string    `json:"-" gorm:"unique;type:varchar(256)"`
	Email                string    `json:"email" gorm:"type:varchar(256)"`
	NormalizedEmail      string    `json:"-" gorm:"uniqueIndex;type:varchar(256)"`
	EmailConfirmed       bool      `json:"emailconfirmed" gorm:"not null"`
	PasswordHash         string    `json:"-" gorm:"type:text"`
	SecurityStamp        string    `json:"-" gorm:"type:text"`
	PhoneNumber          string    `json:"phonenumber" gorm:"type:text"`
	PhoneNumberConfirmed bool      `json:"phonenumberconfirmed" gorm:"not null"`
	TwoFactorEnabled     bool      `json:"twofactorenabled" gorm:"not null"`
	IsTemporaryPassword  bool      `json:"istemporarypassword" gorm:"not null"`
	LockoutEnd           null.Time `json:"lockoutend" gorm:""`
	LockoutEnabled       bool      `json:"lockoutenabled" gorm:"not null"`
	AccessFailedCount    int       `json:"accessfailedcount" gorm:"type:int;not null"`

	Record     `gorm:"embedded"`
	UserClaims []UserClaims `gorm:"foreignKey:UserId"`
	Logins     []UserLogins `gorm:"foreignKey:UserId"`
}

// //TableName 数据表名称
// func (m *User) TableName() string {
// 	return "Users"
// }

func (r *User) BeforeCreate(tx *gorm.DB) error {

	r.NormalizedAccount = strings.ToUpper(r.Account)
	r.NormalizedEmail = strings.ToUpper(r.Email)
	return nil
}
func (r *User) BeforeUpdate(tx *gorm.DB) error {

	r.NormalizedAccount = strings.ToUpper(r.Account)
	r.NormalizedEmail = strings.ToUpper(r.Email)
	return nil
}

// Implement model.IUser interface
func (r *User) GetId() string {
	return r.Id
}
func (r *User) GetUserName() string {
	return r.NormalizedAccount
}
func (r *User) GetEmail() string {
	return r.Email
}
func (r *User) GetPasswordHash() string {
	return r.PasswordHash
}
