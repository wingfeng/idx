package models

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

// Record 数据库通用结构
type Record struct {
	Creator   string         `json:"creator" gorm:"type:varchar(255)"`
	CreatorId string         `json:"creatorid" gorm:"type:varchar(36)"`
	Updator   string         `json:"updator" gorm:"type:varchar(255)"`
	UpdatorId string         `json:"updatorid" gorm:"type:varchar(36)"`
	CreatedAt null.Time      `json:"created" gorm:"autoCreateTime"`
	UpdatedAt null.Time      `json:"updated" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// func (r *Record) BeforeCreate(tx *gorm.DB) error {
// 	u, _ := tx.Get(Const_UserNameKey)
// 	uId, _ := tx.Get(Const_UserIdKey)
// 	r.Creator = fmt.Sprintf("%v", u)
// 	r.CreatorId = fmt.Sprintf("%v", uId)
// 	return nil
// }

// func (r *Record) BeforeUpdate(tx *gorm.DB) error {
// 	u, _ := tx.Get(Const_UserNameKey)
// 	uId, _ := tx.Get(Const_UserIdKey)
// 	r.Updator = fmt.Sprintf("%v", u)
// 	r.UpdatorId = fmt.Sprintf("%v", uId)
// 	return nil
// }
