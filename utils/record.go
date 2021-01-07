package utils

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

//Record 数据库通用结构
type Record struct {
	Creator   string         `json:"creator" gorm:"type:varchar(255)"`
	CreatorID string         `json:"creatorid" gorm:"type:varchar(36)"`
	Updator   string         `json:"updator" gorm:"type:varchar(255)"`
	UpdatorID string         `json:"updatorid" gorm:"type:varchar(36)"`
	CreatedAt null.Time      `json:"created" gorm:"autoCreateTime"`
	UpdatedAt null.Time      `json:"updated" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// func (r *Record) BeforeCreate(tx *gorm.DB) error {
// 	u, _ := tx.Get(Const_UserNameKey)
// 	uID, _ := tx.Get(Const_UserIDKey)
// 	r.Creator = fmt.Sprintf("%v", u)
// 	r.CreatorID = fmt.Sprintf("%v", uID)
// 	return nil
// }

// func (r *Record) BeforeUpdate(tx *gorm.DB) error {
// 	u, _ := tx.Get(Const_UserNameKey)
// 	uID, _ := tx.Get(Const_UserIDKey)
// 	r.Updator = fmt.Sprintf("%v", u)
// 	r.UpdatorID = fmt.Sprintf("%v", uID)
// 	return nil
// }
