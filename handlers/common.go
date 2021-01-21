package handlers

import (
	"crypto/rsa"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	log "github.com/cihub/seelog"
	"github.com/dgrijalva/jwt-go"
	"github.com/wingfeng/idx/oauth2/server"
)

var HTMLTemplate *template.Template
var PublicKey *rsa.PublicKey
var Srv *server.Server

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
func verifyAuthorizationToken(r *http.Request) (jwt.MapClaims, error) {
	header := r.Header.Get("Authorization")
	tokenString := strings.Split(header, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return PublicKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		err = fmt.Errorf("解析Token错误,Validate:%b,Error:%s", token.Valid, err.Error())
		log.Errorf(err.Error())
		return nil, err
	}
}
