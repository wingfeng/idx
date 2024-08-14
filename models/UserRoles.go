package models

// UserRoles [...]
type UserRoles struct {
	UserId string `gorm:"primary_key;type:varchar(255);not null"`

	RoleId string `gorm:"primary_key;index:IX_UserRoles_RoleId;type:varchar(255);not null"`
	OUId   string `gorm:"column:ou_id;primary_key;type:varchar(36);not null"`

	Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *UserRoles) TableName() string {
// 	return "UserRoles"
// }

func (m *UserRoles) GetID() interface{} {
	return m
}
