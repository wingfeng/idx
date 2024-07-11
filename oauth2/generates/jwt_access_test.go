package generates_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wingfeng/idx/oauth2"
	"github.com/wingfeng/idx/oauth2/generates"
	"github.com/wingfeng/idx/oauth2/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJWTAccess(t *testing.T) {
	Convey("Test JWT Access Generate", t, func() {
		data := &oauth2.GenerateBasic{
			Issuer: "http://localhost.idx",
			Client: &models.Client{
				ID:     "123456",
				Secret: "123456",
			},
			UserID: "000000",
			TokenInfo: &models.Token{
				AccessCreateAt:  time.Now(),
				AccessExpiresIn: time.Second * 120,
			},
			Nonce: "Nonce 123456",
		}

		gen := generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512)
		access, refresh, err := gen.Token(context.Background(), data, true)
		So(err, ShouldBeNil)
		So(access, ShouldNotBeEmpty)
		So(refresh, ShouldNotBeEmpty)

		token, err := jwt.ParseWithClaims(access, &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("parse error")
			}
			return []byte("00000000"), nil
		})
		So(err, ShouldBeNil)

		claims, ok := token.Claims.(*generates.JWTAccessClaims)
		So(ok, ShouldBeTrue)
		So(token.Valid, ShouldBeTrue)
		So(claims.Audience, ShouldEqual, "123456")
		So(claims.Subject, ShouldEqual, "000000")
		So(claims.Issuer, ShouldEqual, "http://localhost.idx")
		So(claims.Nonce, ShouldEqual, "Nonce 123456")
	})
}
