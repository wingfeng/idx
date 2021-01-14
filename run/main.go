package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/pem"
	"io/ioutil"
	"net/http"

	log "github.com/cihub/seelog"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
	"github.com/wingfeng/idx/core"
	"github.com/wingfeng/idx/handlers"
	idxmodels "github.com/wingfeng/idx/models"
	idxstore "github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
)

func main() {
	manager := core.NewDefaultManager()
	//	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	tStore, _ := store.NewMemoryTokenStore()
	// token store
	manager.SetTokenStore(tStore)
	privateKeyByets, err := ioutil.ReadFile("../certs/rsa_2048_priv.pem")
	if err != nil {
		log.Errorf("读取私钥错误!,Err:%s", err.Error())
	}

	//pkBlock, _ := pem.Decode(privateKeyByets)
	//privateKey, _ := x509.ParsePKCS1PrivateKey(pkBlock.Bytes)
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
	jwks := &core.JWKS{}
	jwk := core.NewRSAJWTKey()
	jwk.N = base64.URLEncoding.EncodeToString(pi.(*rsa.PublicKey).N.Bytes())
	var buf = make([]byte, 8)
	e := uint64(pi.(*rsa.PublicKey).E)
	binary.LittleEndian.PutUint64(buf, e)
	//base64.URLEncoding.EncodeToString()
	jwk.E = base64.URLEncoding.EncodeToString(buf)

	jwtAccessGenerate := generates.NewJWTAccessGenerate("", privateKeyByets, jwt.SigningMethodRS256)
	jwk.Alg = jwtAccessGenerate.SignedMethod.Alg()
	jwk.Kid = []byte(jwtAccessGenerate.SignedKeyID)

	jwks.Keys = append(jwks.Keys, jwk)

	handlers.Jwks = jwks
	// generate jwt access token
	manager.MapAccessGenerate(jwtAccessGenerate)

	//初始化DB
	db := utils.GetDB("mysql", "root:123456@tcp(localhost:3306)/sso?&parseTime=true")
	idxmodels.Sync2Db(db)
	clientStore := idxstore.NewClientStore(db)
	userStore := idxstore.NewDbUserStore(db)

	handlers.UserStore = userStore
	manager.SetClientStore(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})
	srv.SetClientScopeHandler(func(clientid, scope string) (allow bool, err error) {
		return true, nil
	})
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Infof("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Infof("Response Error:", re.Error.Error())
	})
	handlers.Srv = srv

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/auth", handlers.AuthHandler)

	http.HandleFunc("/connect/authorize", handlers.Authorize)

	http.HandleFunc("/connect/token", handlers.Token)

	http.HandleFunc("/test", handlers.Test)
	http.HandleFunc("/.well-known/openid-configuration", handlers.WellknownHandler)
	http.HandleFunc("/.well-known/openid-configuration/jwks", handlers.JWKSHandler)
	log.Infof("Server is running at 9096 port.")
	log.Error(http.ListenAndServe(":9096", nil))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}
		returnURI := r.Form
		store.Set("ReturnUri", returnURI)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	//	store.Delete("LoggedInUserID")
	store.Save()
	return
}
