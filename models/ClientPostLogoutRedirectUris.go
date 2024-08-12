package models

// ClientPostLogoutRedirectURIs [...]
type ClientPostLogoutRedirectURIs struct {
	Id                    int    `gorm:"primary_key;auto_Increment;not null"`
	PostLogoutRedirectURI string `gorm:"type:varchar(2000);not null"`
	ClientId              int    `gorm:"type:varchar(256);not null"`

	Record `gorm:"embedded"`
}
