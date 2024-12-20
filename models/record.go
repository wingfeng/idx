package models

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

const (
	CONST_USERNAMEKEY = "user"
	CONST_USERIDKEY   = "userid"
	CONST_OUKEY       = "org"
	CONST_OUIDKEY     = "orgid"
)

// Record 数据库通用结构
type Record struct {
	Creator   string    `json:"creator" gorm:"type:varchar(255)"`
	CreatorId string    `json:"creatorId" gorm:"type:varchar(36)"`
	Updator   string    `json:"updator" gorm:"type:varchar(255)"`
	UpdatorId string    `json:"updatorId" gorm:"type:varchar(36)"`
	CreatedAt null.Time `json:"created" gorm:"autoCreateTime"`
	UpdatedAt null.Time `json:"updated" gorm:"autoUpdateTime"`
	//	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type IntRecord struct {
	Id int64 `json:"id,omitempty" gorm:"primary_key;auto_Increment;not null"`
	Record
}

func (r *IntRecord) GetID() interface{} {
	return r.Id
}

// SnowflakeRecord 数据库通用结构
// ID generated by snowflake
type SnowflakeRecord struct {
	Id snowflake.ID `json:"id,omitempty" gorm:"primary_key;not null"`
	Record
}

func (r *SnowflakeRecord) GetID() interface{} {
	return r.Id
}

func (r *Record) BeforeCreate(tx *gorm.DB) error {
	u, _ := tx.Get(CONST_USERNAMEKEY)
	uId, _ := tx.Get(CONST_USERIDKEY)
	r.Creator = fmt.Sprintf("%v", u)
	r.CreatorId = fmt.Sprintf("%v", uId)
	return nil
}

func (r *Record) BeforeUpdate(tx *gorm.DB) error {
	u, _ := tx.Get(CONST_USERNAMEKEY)
	uId, _ := tx.Get(CONST_USERIDKEY)
	r.Updator = fmt.Sprintf("%v", u)
	r.UpdatorId = fmt.Sprintf("%v", uId)
	return nil
}
