package models

import (
	"strings"

	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

// User 用户信息
type User struct {
	ID                   string    `json:"id" gorm:"primary_key;column:Id;type:varchar(255);not null"`
	OUID                 string    `json:"ouid" gorm:"column:OUId;type:varchar(36)"`
	OU                   string    `json:"ou" gorm:"column:OU;type:varchar(256)"`
	Account              string    `json:"account" gorm:"column:Account;type:varchar(256)"`
	DisplayName          string    `json:"displayname" gorm:"column:DisplayName;type:varchar(256)"`
	NormalizedAccount    string    `json:"-" gorm:"unique;column:NormalizedAccount;type:varchar(256)"`
	Email                string    `json:"email" gorm:"column:Email;type:varchar(256)"`
	NormalizedEmail      string    `json:"-" gorm:"index:EmailIndex;column:NormalizedEmail;type:varchar(256)"`
	EmailConfirmed       bool      `json:"emailconfirmed" gorm:"column:EmailConfirmed;type:tinyint(1);not null"`
	PasswordHash         string    `json:"-" gorm:"column:PasswordHash;type:longtext"`
	SecurityStamp        string    `json:"-" gorm:"column:SecurityStamp;type:longtext"`
	ConcurrencyStamp     string    `json:"-" gorm:"column:ConcurrencyStamp;type:longtext"`
	PhoneNumber          string    `json:"phonenumber" gorm:"column:PhoneNumber;type:longtext"`
	PhoneNumberConfirmed bool      `json:"phonenumberconfirmed" gorm:"column:PhoneNumberConfirmed;type:tinyint(1);not null"`
	TwoFactorEnabled     bool      `json:"twofactorenabled" gorm:"column:TwoFactorEnabled;type:tinyint(1);not null"`
	IsTemporaryPassword  bool      `json:"istemporarypassword" gorm:"column:IsTemporaryPassword;type:tinyint(1);not null"`
	LockoutEnd           null.Time `json:"lockoutend" gorm:"column:LockoutEnd;type:datetime(6)"`
	LockoutEnabled       bool      `json:"lockoutenabled" gorm:"column:LockoutEnabled;type:tinyint(1);not null"`
	AccessFailedCount    int       `json:"accessfailedcount" gorm:"column:AccessFailedCount;type:int(11);not null"`

	utils.Record `gorm:"embedded"`
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
