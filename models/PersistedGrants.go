package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// PersistedGrants [...]
type PersistedGrants struct {
	Key          string    `gorm:"primary_key;column:Key;type:varchar(200);not null"`
	Type         string    `gorm:"index:IX_PersistedGrants_SubjectId_ClientId_Type;column:Type;type:varchar(50);not null"`
	SubjectID    string    `gorm:"index:IX_PersistedGrants_SubjectId_ClientId_Type;column:SubjectId;type:varchar(200)"`
	ClientID     string    `gorm:"index:IX_PersistedGrants_SubjectId_ClientId_Type;column:ClientId;type:varchar(200);not null"`
	CreationTime time.Time `gorm:"column:CreationTime;type:datetime(6);not null"`
	Expiration   null.Time `gorm:"index:IX_PersistedGrants_Expiration;column:Expiration;type:datetime(6)"`
	Data         string    `gorm:"column:Data;type:longtext;not null"`
	utils.Record `gorm:"embedded"`
}
