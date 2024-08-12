package models

// RoleClaims [...]
type RoleClaims struct {
	Id         int    `gorm:"primary_key;auto_Increment;not null"`
	RoleId     string `gorm:"index:IX_RoleClaims_RoleId;type:varchar(255);not null"`
	ClaimType  string `gorm:"type:text"`
	ClaimValue string `gorm:"type:text"`
	Record     `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *RoleClaims) TableName() string {
// 	return "RoleClaims"
// }
