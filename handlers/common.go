package handlers

import (
	"crypto/rsa"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	log "github.com/cihub/seelog"
	"github.com/golang-jwt/jwt"
	"github.com/wingfeng/idx/oauth2/server"
	"github.com/wingfeng/idx/store"
)

var HTMLTemplate *template.Template
var PublicKey *rsa.PublicKey
var Srv *server.Server
var ClientStore *store.ClientStore

func verifyAuthorizationToken(r *http.Request) (jwt.MapClaims, error) {
	header := r.Header.Get("Authorization")
	tokenString := strings.Split(header, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return PublicKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		err = fmt.Errorf("解析Token错误,Validate:%v,Error:%s", token.Valid, err.Error())
		log.Errorf(err.Error())
		return nil, err
	}
}
