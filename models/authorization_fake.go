package models

//Ony use for define gorm db schema
type Authorization_fake struct {
	Id                    string `json:"id" gorm:"primary_key;type:varchar(36);not null"`
	Issuer                string `json:"iss" gorm:"type:varchar(256);"`
	ClientId              string `json:"client_id" gorm:"type:varchar(256);"`
	PrincipalName         string `json:"principal_name" gorm:"type:varchar(256);"`
	Subject               string `json:"subject" gorm:"type:varchar(256);"`
	GrantType             string `json:"grant_type" gorm:"type:varchar(256);"`
	ResponseType          string `json:"response_type" gorm:"type:varchar(256);"`
	Scope                 string `json:"scope" gorm:"type:varchar(256);"`
	Code                  string `json:"code" gorm:"type:varchar(256);index"`
	Nonce                 string `json:"nonce" gorm:"type:varchar(256);"`
	AccessToken           string `json:"access_token" gorm:"type:varchar(2560);index" `
	RefreshToken          string `json:"refresh_token" gorm:"type:varchar(256);index"`
	IDToken               string `json:"id_token" gorm:"type:varchar(2560);index"`
	CodeChallenge         string `json:"code_challenge" gorm:"type:varchar(256);"`
	CodeChallengeMethod   string `json:"code_challenge_method" gorm:"type:varchar(256);"`
	DeviceCode            string `json:"device_code" gorm:"type:varchar(256);index"`
	ExpiresAt             int64  `json:"expires_at"` //Unix timestamp
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
	UserCode              string `json:"user_code" gorm:"type:varchar(256);index"`
}

func (r *Authorization_fake) TableName() string {
	return "authorizations"
}
