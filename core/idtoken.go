package core

// {
// 	"iss": "http://server.example.com",
// 	"sub": "248289761001",
// 	"aud": "s6BhdRkqt3",
// 	"nonce": "n-0S6_WzA2Mj",
// 	"exp": 1311281970,
// 	"iat": 1311280970
//   }
type IDToken struct {
	Issuer  string `json:"iss"`
	Sub     string `json:"sub"`
	Aud     string `json:"aud"`
	Nonce   string `json:"nonce"`
	Expire  int32  `json:"exp"`
	IssueAt int32  `json:"iat"`
}
