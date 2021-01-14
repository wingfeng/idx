package models

import "github.com/wingfeng/idx/utils"

// ClientIDPRestrictions [...]
type ClientIDPRestrictions struct {
	ID           int    `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Provider     string `gorm:"column:Provider;type:varchar(200);not null"`
	ClientID     int    `gorm:"index:IX_ClientIdPRestrictions_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Client `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}

//TableName 数据表名称
func (m *ClientIDPRestrictions) TableName() string {
	return "client_idp_restrictions"
}
