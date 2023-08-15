package models

import "github.com/wingfeng/idx/utils"

// ClientIDPRestrictions [...]
type ClientIDPRestrictions struct {
	ID           int    `gorm:"primary_key;auto_Increment;column:id;not null"`
	Provider     string `gorm:"column:provider;type:varchar(200);not null"`
	ClientID     int    `gorm:"index:IX_ClientIdPRestrictions_ClientId;column:clientid;type:int;not null"`
	Clients      Client `gorm:"foreignKey:clientid;references:id"`
	utils.Record `gorm:"embedded"`
}

// TableName 数据表名称
func (m *ClientIDPRestrictions) TableName() string {
	return "client_idp_restrictions"
}
