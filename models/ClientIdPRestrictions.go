package models

import "github.com/wingfeng/idx/utils"

// ClientIDPRestrictions [...]
type ClientIDPRestrictions struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:Id;not null"`
	Provider     string `gorm:"column:Provider;type:varchar(200);not null"`
	ClientID     int    `gorm:"index:IX_ClientIdPRestrictions_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Client `gorm:"foreignkey:ClientId"`
	utils.Record `gorm:"embedded"`
}

//TableName 数据表名称
func (m *ClientIDPRestrictions) TableName() string {
	return "client_idp_restrictions"
}
