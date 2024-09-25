package models

// ClientPostLogoutRedirectURIs [...]
type ClientPostLogoutRedirectURIs struct {
	PostLogoutRedirectURI string `gorm:"type:varchar(2000);not null"`
	ClientId              int64  `gorm:"type:bigint;not null"`

	IntRecord `gorm:"embedded"`
}
