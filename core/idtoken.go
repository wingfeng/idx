package core

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// {
// 	"iss": "http://server.example.com",
// 	"sub": "248289761001",
// 	"aud": "s6BhdRkqt3",
// 	"nonce": "n-0S6_WzA2Mj",
// 	"exp": 1311281970,
// 	"iat": 1311280970
//   }
type IDToken struct {
	Issuer          string `json:"iss,omitempty"`
	Sub             string `json:"sub,omitempty"`
	Aud             string `json:"aud,omitempty"`
	Nonce           string `json:"nonce,omitempty"`
	Expire          int64  `json:"exp,omitempty"`
	AccessTokenHash string `json:"at_hash"`
	IssueAt         int64  `json:"iat,omitempty"`
}

func (idt *IDToken) GetClaims() jwt.MapClaims {
	result := jwt.MapClaims{}
	result["iss"] = idt.Issuer
	result["sub"] = idt.Sub
	result["aud"] = idt.Aud
	if !strings.EqualFold(idt.Nonce, "") {
		result["nonce"] = idt.Nonce
	}
	result["exp"] = idt.Expire
	result["iat"] = idt.IssueAt
	result["at_hash"] = idt.AccessTokenHash
	return result
}
