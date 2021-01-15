package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/cihub/seelog"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/spf13/viper"
	"github.com/wingfeng/idx/core"
	"github.com/wingfeng/idx/handlers"
	"github.com/wingfeng/idx/models"
	idxmodels "github.com/wingfeng/idx/models"
	idxstore "github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
)

func main() {
	option := initConfig()

	if option.SyncDB {
		//初始化DB
		dbEngine := utils.GetDB(option.Driver, option.Connection)
		models.Sync2Db(dbEngine)
		fmt.Println("同步数据库结构完成")
		return
	}
	manager := core.NewDefaultManager()
	//	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	tStore, _ := store.NewMemoryTokenStore()
	// token store
	manager.SetTokenStore(tStore)
	privateKeyByets, err := ioutil.ReadFile(option.PrivateKeyPath)
	if err != nil {
		log.Errorf("读取私钥错误!,Err:%s", err.Error())
	}

	publicKeyBytes, err := ioutil.ReadFile(option.PublicKeyPath)
	if err != nil {
		log.Errorf("读取公钥错误!,Err:%s", err.Error())
	}

	jwks := &core.JWKS{}
	jwk := core.NewRSAJWTKeyWithPEM(publicKeyBytes)

	jwtAccessGenerate := generates.NewJWTAccessGenerate("", privateKeyByets, jwt.SigningMethodRS256)
	jwk.Alg = jwtAccessGenerate.SignedMethod.Alg()
	jwk.Kid = jwtAccessGenerate.SignedKeyID
	jwks.Keys = append(jwks.Keys, jwk)
	handlers.PublicKey = jwk.PublicKey
	handlers.Jwks = jwks
	// generate jwt access token
	manager.MapAccessGenerate(jwtAccessGenerate)

	//初始化DB
	db := utils.GetDB(option.Driver, option.Connection)
	idxmodels.Sync2Db(db)
	clientStore := idxstore.NewClientStore(db)
	userStore := idxstore.NewDbUserStore(db)

	handlers.UserStore = userStore
	manager.SetClientStore(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	openidExt := core.NewOpenIDExtend()
	openidExt.PrivateKeyByets = privateKeyByets
	openidExt.ClientStore = clientStore
	openidExt.UserStore = userStore

	srv.SetPasswordAuthorizationHandler(openidExt.PasswordAuthorizationHandler)
	srv.SetClientScopeHandler(openidExt.ClientScopeHandler)

	srv.SetUserAuthorizationHandler(openidExt.UserAuthorizeHandler)
	srv.SetExtensionFieldsHandler(openidExt.Id_TokenHandler)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Infof("OAuth Server Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Infof("Response Error:", re.Error.Error())
	})
	// htmlTplEngine := template.New("htmlTplEngine")
	// // 模板根目录下的模板文件 一些公共文件
	// _, htmlTplEngineErr := htmlTplEngine.ParseGlob("../static/*.html")
	// if nil != htmlTplEngineErr {
	// 	seelog.Errorf("解析html模板错误,Error:%s", htmlTplEngineErr.Error())
	// }
	handlers.Srv = srv
	//handlers.HTMLTemplate = htmlTplEngine

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/auth", handlers.AuthHandler)

	http.HandleFunc("/connect/authorize", handlers.Authorize)

	http.HandleFunc("/connect/token", handlers.Token)
	http.HandleFunc("/connect/userinfo", handlers.UserInfoHandler)

	http.HandleFunc("/test", handlers.Test)
	http.HandleFunc("/.well-known/openid-configuration", handlers.WellknownHandler)
	http.HandleFunc("/.well-known/openid-configuration/jwks", handlers.JWKSHandler)
	http.HandleFunc("/connect/endsession", handlers.LogoutHandler)
	http.HandleFunc("/connect/revocation", handlers.RevocateHandler)
	log.Infof("Server is running at 9096 port.")
	log.Error(http.ListenAndServe(fmt.Sprintf(":%d", option.Port), nil))
}
func initConfig() *Option {
	confPath := flag.String("conf", "../conf/config.yaml", "配置文件路径")
	syncDb := flag.Bool("syncdb", false, "同步数据结构到数据库.")
	flag.Parse()

	viper.SetConfigFile(*confPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AllowEmptyEnv(true)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("读取配置文件错误: %s ", err.Error()))
	}
	viper.SetEnvPrefix("IDX")
	viper.AutomaticEnv()

	opts := &Option{}
	opts.SyncDB = *syncDb
	err = viper.Unmarshal(opts)
	if err != nil {
		log.Error("读取配置错误:", err)
	}

	return opts
}
