package test

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	gormstore "github.com/go-session/gorm"
	"github.com/go-session/session"
	"github.com/labstack/gommon/log"
	"github.com/wingfeng/idx/core"
	"github.com/wingfeng/idx/handlers"
	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/oauth2/errors"
	"github.com/wingfeng/idx/oauth2/generates"
	"github.com/wingfeng/idx/oauth2/server"
	"github.com/wingfeng/idx/oauth2/store"
	idxstore "github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
)

func init_router() *gin.Engine {

	dbDriver := "mysql"
	dbConnection := "root:eATq1GDhsP@tcp(localhost:3306)/idx?&parseTime=true"
	certKeyPath := "../certs/rsa_pri.key"
	certPath := "../certs/rsa_pub.pem"
	// hashKey := []byte("FF51A553-72FC-478B-9AEF-93D6F506DE91")
	// session.InitManager(
	// 	session.SetStore(
	// 		cookie.NewCookieStore(
	// 			cookie.SetCookieName("demo_cookie_store_id"),
	// 			cookie.SetHashKey(hashKey),
	// 		),
	// 	),
	// )
	sessionstore := gormstore.MustStore(gormstore.Config{}, dbDriver, dbConnection)

	session.InitManager(
		session.SetStore(sessionstore),
	)
	manager := core.NewDefaultManager()
	//	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	tStore, _ := store.NewMemoryTokenStore()
	// token store
	manager.SetTokenStore(tStore)
	privateKeyBytes, err := ioutil.ReadFile(certKeyPath)
	if err != nil {
		log.Errorf("读取私钥错误!,Err:%s", err.Error())
	}

	publicKeyBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Errorf("读取公钥错误!,Err:%s", err.Error())
	}

	jwks := &core.JWKS{}
	jwk := core.NewRSAJWTKeyWithPEM(publicKeyBytes)
	kid := "d2a820a8916647f7ac72627ec0ae4f94"

	jwtAccessGenerate := generates.NewJWTAccessGenerate(kid, privateKeyBytes, jwt.SigningMethodRS256)
	jwk.Alg = jwtAccessGenerate.SignedMethod.Alg()
	jwk.Kid = jwtAccessGenerate.SignedKeyID
	jwks.Keys = append(jwks.Keys, jwk)
	handlers.PublicKey = jwk.PublicKey
	handlers.Jwks = jwks
	// generate jwt access token
	manager.MapAccessGenerate(jwtAccessGenerate)
	manager.PrivateKeyBytes = privateKeyBytes
	manager.Kid = kid
	//初始化DB
	db := utils.GetDB(dbDriver, dbConnection)
	idxmodels.Sync2Db(db)
	clientStore := idxstore.NewClientStore(db)
	userStore := idxstore.NewDbUserStore(db)

	handlers.ClientStore = clientStore

	manager.SetClientStore(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	openidExt := core.NewOpenIDExtend()
	openidExt.PrivateKeyByets = privateKeyBytes
	openidExt.ClientStore = clientStore
	openidExt.UserStore = userStore

	srv.SetPasswordAuthorizationHandler(openidExt.PasswordAuthorizationHandler)

	srv.Config.AllowedResponseTypes = append(srv.Config.AllowedResponseTypes, "id_token")
	srv.SetUserAuthorizationHandler(openidExt.UserAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Infof("OAuth Server Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Infof("Response Error:", re.Error.Error())
	})

	handlers.Srv = srv
	router := gin.Default()
	loginCtrl := &handlers.LoginController{
		UserStore: *userStore,
	}

	router.GET("/login", loginCtrl.LoginGet)
	router.POST("/login", loginCtrl.LoginPost)
	userCtrl := &handlers.UserInfoController{
		UserStore: userStore,
	}
	router.GET("/connect/userinfo", userCtrl.UserInfo)
	router.GET("/consent", handlers.Consent)

	router.GET("/connect/authorize", handlers.Authorize)
	router.POST("/connect/authorize", handlers.Authorize)

	router.POST("/connect/token", handlers.TokenController)

	router.GET("/test", handlers.Test)
	router.GET("/.well-known/openid-configuration", handlers.WellknownHandler)
	router.GET("/.well-known/openid-configuration/jwks", handlers.JWKSHandler)
	router.POST("/connect/endsession", handlers.LogoutHandler)
	router.POST("/connect/revocation", handlers.RevocateHandler)
	router.LoadHTMLGlob("../static/*")

	return router
	// handler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// 	AllowedHeaders:   []string{"*"},
	// 	// Enable Debugging for testing, consider disabling in production
	// 	Debug: false,
	// }).Handler(router)
}
