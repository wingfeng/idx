package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	log "github.com/cihub/seelog"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	gormstore "github.com/go-session/gorm"
	"github.com/go-session/session"

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
	//配置Log
	consoleWriter, _ := log.NewConsoleWriter() //创建一个新的控制台写入器
	logLevel, lex := log.LogLevelFromString(option.LogLevel)
	if !lex {
		logLevel = log.DebugLvl
	}
	logger, _ := log.LoggerFromWriterWithMinLevel(consoleWriter, logLevel)
	log.ReplaceLogger(logger)
	defer log.Flush()

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

	manager.HTTPScheme = option.HTTPScheme
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
	//clientStore.Cache = rdb
	userStore := idxstore.NewDbUserStore(db)

	handlers.ClientStore = clientStore

	manager.SetClientStore(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	openidExt := core.NewOpenIDExtend()
	openidExt.PrivateKeyByets = privateKeyBytes
	openidExt.ClientStore = clientStore
	openidExt.UserStore = userStore
	manager.UserStore = userStore

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
	router.POST("/connect/userinfo", userCtrl.UserInfo)
	router.GET("/consent", handlers.Consent)

	router.GET("/connect/authorize", handlers.Authorize)
	router.POST("/connect/authorize", handlers.Authorize)

	router.POST("/connect/token", handlers.TokenController)

	router.GET("/test", handlers.Test)
	wellknowCtrl := &handlers.WellknownController{
		Scheme: option.HTTPScheme,
	}
	router.GET("/.well-known/openid-configuration", wellknowCtrl.Get)
	router.GET("/.well-known/openid-configuration/jwks", handlers.JWKSHandler)
	router.POST("/connect/endsession", handlers.LogoutHandler)
	router.POST("/connect/revocation", handlers.RevocateHandler)
	router.LoadHTMLGlob("../static/*")
	log.Infof("Server is running at %d port.", option.Port)
	// handler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// 	AllowedHeaders:   []string{"*"},
	// 	// Enable Debugging for testing, consider disabling in production
	// 	Debug: false,
	// }).Handler(router)
	//	handler := cors.Default().Handler(mux)
	address := fmt.Sprintf("%s:%d", "", option.Port)
	//l := logger{}
	//	router.RunTLS(address, "../certs/ca/localhost/localhost.crt", "../certs/ca/localhost/localhost.key")
	router.Run(address)
	//err = http.ListenAndServe(address, handler) //accesslog.NewLoggingHandler(handler, l))
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
