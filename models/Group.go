package models

type Group struct {
	SnowflakeRecord
	Name        string `json:"name" gorm:"type:varchar(200);not null"`
	Description string `json:"description" gorm:"type:varchar(1000)"`
	Members     []User `gorm:"many2many:user_groups;"`
}
