package models

// PersistedGrants [...]
type PersistedGrants struct {
	Principal string `gorm:"primaryKey;type:varchar(200)"`
	ClientId  string `gorm:"primaryKey;type:varchar(200);not null"`
	Scope     string `gorm:"type:text;not null"`
	Record    `gorm:"embedded"`
}

func (m *PersistedGrants) TableName() string {
	return "persisted_grants"
}
