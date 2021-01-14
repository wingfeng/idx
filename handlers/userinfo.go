package handlers

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/cihub/seelog"
	"github.com/dgrijalva/jwt-go"
)

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	tokenString := strings.Split(header, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		publicKeyBytes, err := ioutil.ReadFile("../certs/rsa_2048_pub.pem")
		if err != nil {
			log.Errorf("读取公钥错误!,Err:%s", err.Error())
		}
		block, _ := pem.Decode(publicKeyBytes)
		if block == nil {
			log.Error("public key error")
		}
		// 解析公钥
		pi, err := x509.ParsePKIXPublicKey(block.Bytes)

		return pi, err
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["sub"]
		user, err := UserStore.GetUserByID(id.(string))
		result := make(map[string]interface{})
		result["sub"] = user.ID
		result["email"] = user.Email
		result["email_verified"] = user.EmailConfirmed
		result["display_name"] = user.DisplayName
		result["ou"] = user.OU
		result["ouid"] = user.OUID
		if err != nil {
			log.Errorf("获取用户错误,Error:%s", token.Valid, err.Error())
		}
		json.NewEncoder(w).Encode(result)
	} else {
		log.Errorf("解析Token错误,Validate:%b,Error:%s", token.Valid, err.Error())
	}
}
