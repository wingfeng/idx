package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
)

// AuditLog [...]
type AuditLog struct {
	ID                    int       `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Event                 string    `gorm:"column:Event;type:longtext"`
	Source                string    `gorm:"column:Source;type:longtext"`
	Category              string    `gorm:"column:Category;type:longtext"`
	SubjectIDentifier     string    `gorm:"column:SubjectIdentifier;type:longtext"`
	SubjectName           string    `gorm:"column:SubjectName;type:longtext"`
	SubjectType           string    `gorm:"column:SubjectType;type:longtext"`
	SubjectAdditionalData string    `gorm:"column:SubjectAdditionalData;type:longtext"`
	Action                string    `gorm:"column:Action;type:longtext"`
	Data                  string    `gorm:"column:Data;type:longtext"`
	Created               time.Time `gorm:"column:Created;type:datetime(6);not null"`
	utils.Record          `gorm:"embedded"`
}
