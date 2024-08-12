package models

// ClientIdPRestrictions [...]
type ClientIdPRestrictions struct {
	Id       int    `gorm:"primary_key;auto_Increment;not null"`
	Provider string `gorm:"type:varchar(200);not null"`
	ClientId int    `gorm:"index:IX_ClientIdPRestrictions_ClientId;type:int;not null"`

	Record `gorm:"embedded"`
}

// TableName 数据表名称
func (m *ClientIdPRestrictions) TableName() string {
	return "client_idp_restrictions"
}
