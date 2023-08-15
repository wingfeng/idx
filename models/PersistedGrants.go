package models

import (
	"time"

	"github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
)

// PersistedGrants [...]
type PersistedGrants struct {
	Key          string    `gorm:"primary_key;column:Key;type:varchar(200);not null"`
	Type         string    `gorm:"index:IX_PersistedGrants_SubjectId_ClientId_Type;column:type;type:varchar(50);not null"`
	SubjectID    string    `gorm:"index:IX_PersistedGrants_SubjectId_ClientId_Type;column:subjectid;type:varchar(200)"`
	ClientID     string    `gorm:"index:IX_PersistedGrants_SubjectId_ClientId_Type;column:clientid;type:varchar(200);not null"`
	CreationTime time.Time `gorm:"column:creationtime;not null"`
	Expiration   null.Time `gorm:"index:IX_PersistedGrants_Expiration;column:expiration;"`
	Data         string    `gorm:"column:data;type:text;not null"`
	utils.Record `gorm:"embedded"`
}
