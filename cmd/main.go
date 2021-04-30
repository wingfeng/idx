package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/dgrijalva/jwt-go"
	gormstore "github.com/go-session/gorm"
	"github.com/go-session/session"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/wingfeng/idx/core"
	"github.com/wingfeng/idx/handlers"
	"github.com/wingfeng/idx/models"
	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/oauth2/errors"
	"github.com/wingfeng/idx/oauth2/generates"
	"github.com/wingfeng/idx/oauth2/server"
	"github.com/wingfeng/idx/oauth2/store"
	idxstore "github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
)

var (
	hashKey = []byte("FF51A553-72FC-478B-9AEF-93D6F506DE91")

	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
)

func main() {
	showVersion := flag.Bool("ver", false, "程序版本")
	flag.Parse()
	if *showVersion {
		Version()
		return
	}
	option := initConfig()
	sessionstore := gormstore.MustStore(gormstore.Config{}, option.Driver, option.Connection)
	defer sessionstore.Close()

	session.InitManager(
		session.SetStore(sessionstore),
	)

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
	privateKeyBytes, err := ioutil.ReadFile(option.PrivateKeyPath)
	if err != nil {
		log.Errorf("读取私钥错误!,Err:%s", err.Error())
	}

	publicKeyBytes, err := ioutil.ReadFile(option.PublicKeyPath)
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
	db := utils.GetDB(option.Driver, option.Connection)
	idxmodels.Sync2Db(db)
	clientStore := idxstore.NewClientStore(db)
	userStore := idxstore.NewDbUserStore(db)

	handlers.ClientStore = clientStore
	handlers.UserStore = userStore

	manager.SetClientStore(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	openidExt := core.NewOpenIDExtend()
	openidExt.PrivateKeyByets = privateKeyBytes
	openidExt.ClientStore = clientStore
	openidExt.UserStore = userStore

	srv.SetPasswordAuthorizationHandler(openidExt.PasswordAuthorizationHandler)
	//	srv.SetClientScopeHandler(openidExt.ClientScopeHandler)
	srv.Config.AllowedResponseTypes = append(srv.Config.AllowedResponseTypes, "id_token")
	srv.SetUserAuthorizationHandler(openidExt.UserAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Infof("OAuth Server Internal Error:", err.Error())
		return
	})
	// //	srv.ClientInfoHandler = server.ClientFormHandler
	// srv.SetClientInfoHandler(server.ClientFormHandler)
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
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/auth", handlers.AuthHandler)

	mux.HandleFunc("/connect/authorize", handlers.Authorize)

	mux.HandleFunc("/connect/token", handlers.Token)
	mux.HandleFunc("/connect/userinfo", handlers.UserInfoHandler)

	mux.HandleFunc("/test", handlers.Test)
	mux.HandleFunc("/.well-known/openid-configuration", handlers.WellknownHandler)
	mux.HandleFunc("/.well-known/openid-configuration/jwks", handlers.JWKSHandler)
	mux.HandleFunc("/connect/endsession", handlers.LogoutHandler)
	mux.HandleFunc("/connect/revocation", handlers.RevocateHandler)
	log.Infof("Server is running at %d port.", option.Port)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	}).Handler(mux)
	//	handler := cors.Default().Handler(mux)
	address := fmt.Sprintf("%s:%d", "", option.Port)
	err = http.ListenAndServe(address, handler)
	if err != nil {
		log.Error("Server Error:%s", err.Error())
	}

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
func Version() {
	fmt.Printf("App Name:\t%s\n", AppName)
	fmt.Printf("App Version:\t%s\n", AppVersion)
	fmt.Printf("Build version:\t%s\n", BuildVersion)
	fmt.Printf("Build time:\t%s\n", BuildTime)
	fmt.Printf("Git revision:\t%s\n", GitRevision)
	fmt.Printf("Git branch:\t%s\n", GitBranch)
	fmt.Printf("Golang Version: %s\n", GoVersion)
}
