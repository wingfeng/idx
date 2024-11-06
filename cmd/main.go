package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	oauth2 "github.com/wingfeng/idx-oauth2"
	"github.com/wingfeng/idx-oauth2/conf"
	"github.com/wingfeng/idx-oauth2/service"
	"github.com/wingfeng/idx-oauth2/service/impl"
	myConf "github.com/wingfeng/idx/conf"
	"github.com/wingfeng/idx/controller"
	"github.com/wingfeng/idx/ldap"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/models/dto"
	"github.com/wingfeng/idx/repo"
	myService "github.com/wingfeng/idx/service"
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

	syncDb := flag.Bool("syncdb", false, "同步数据结构到数据库.")
	// dbDriver := *flag.String("db", "pgx", "DB Driver:mysql,pgx")
	// dbConnection := *flag.String("dbConnection", "host=localhost port=5432 user=root password=pass@word1 dbname=idx sslmode=disable TimeZone=Asia/Shanghai", "DB Connection")
	// port := flag.Int("port", 9097, "Server Port")

	flag.Parse()
	if *showVersion {
		Version()
		return
	}
	option := myConf.Options
	redisLink := fmt.Sprintf("%s:%d", option.RedisHost, option.RedisPort)

	if *syncDb {
		//初始化DB
		dbEngine := utils.GetDB(option.Driver, option.Connection)
		models.Sync2Db(dbEngine)
		fmt.Println("同步数据库结构完成,程序退出")
		return
	}

	//配置Log
	logLevel := slog.LevelWarn
	switch strings.ToLower(option.LogLevel) {
	case "debug":
		logLevel = slog.LevelDebug

	case "info":
		logLevel = slog.LevelInfo

	case "warn":
		logLevel = slog.LevelWarn

	}
	slog.Info("Set log level", "Level", logLevel)
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(handler))

	config := conf.DefaultConfig()

	//初始化DB

	scopeRepo := repo.NewScopeRepository()
	scopes, err := scopeRepo.GetSupportScopes()
	if err != nil {
		panic(err)
	}
	config.ScopesSupported = scopes
	router := gin.Default()
	store, err := redis.NewStore(10, "tcp", redisLink, "", []byte("secret"))
	if err != nil {
		slog.Error("Redis NewStore Error", "error", err)
		panic(err)
	}
	authRepo := repo.NewAuthorizationRepository()
	userRepo := repo.NewUserRepository()
	consentRepo := repo.NewConsentRepository()
	clientRepo := repo.NewClientRepository()
	tokenService, jwks := buildTokenService(config, userRepo)
	us := myService.NewUserService(userRepo)
	tenant := oauth2.NewTenant(config,
		clientRepo,
		authRepo,
		consentRepo,
		us,
		tokenService, jwks)
	router.LoadHTMLGlob("../static/*.html")
	router.Static("/img", "../static/img")
	authCtrl := controller.NewAuthController(us)
	tenant.LoginCtrl = authCtrl
	tenant.LogoutCtrl = authCtrl
	tenant.InitOAuth2Router(router, sessions.Sessions("idx_session", store))
	//	authCtrl := controller.NewAuthController(userRepo)
	g := router.Group("idx")
	authCtrl.RegistRoute(g)
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/idx")
	})

	go func() {
		address := fmt.Sprintf("%s:%d", "", option.Port)
		slog.Info("Server is running at", "port", option.Port)

		//	router.RunTLS(address, "../certs/ca/localhost/localhost.crt", "../certs/ca/localhost/localhost.key")
		err = router.Run(address)

		if err != nil {
			slog.Error("Server Error", "error", err)
		}

	}()

	//start LDAP Server
	ldap.StartLdapServer(us)

}
func buildTokenService(config *conf.Config, userRepo *repo.DBUserRepository) (service.TokenService, *conf.JWKS) {
	//load private key from pem file
	var privateKey *rsa.PrivateKey
	privateKeyPEM, _ := os.ReadFile(myConf.Options.PrivateKeyPath)

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		slog.Error("failed to decode PEM block containing private key\n gernerating a new private key")
		privateKey = generatePrivateKey(myConf.Options.PrivateKeyPath)
	} else {
		privateKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
	}
	publicKey := &privateKey.PublicKey

	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(publicKey)
	// Convert the RSA public key to PEM format.
	pemPublicKey := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyPEM := pem.EncodeToMemory(pemPublicKey)

	key := conf.NewRSAJWTKeyWithPEM(publicKeyPEM)
	key.Use = "sig"
	key.Kid = "d2a820a8916647f7ac72627ec0ae4f94"
	key.Alg = "RS256"
	jwks := &conf.JWKS{Keys: []interface{}{key}}

	tokenService := impl.NewJwtTokenService(jwt.SigningMethodRS256, privateKey, func(token *jwt.Token, userName string, scope string) map[string]interface{} {
		u, _ := userRepo.GetUserByName(userName)
		token.Header["kid"] = "d2a820a8916647f7ac72627ec0ae4f94"
		user := u.(*dto.UserDto)
		result := map[string]interface{}{}
		//
		if strings.Contains(scope, "mobile") {
			result["mobile"] = user.PhoneNumber
		}
		if strings.Contains(scope, "email") {
			result["email"] = user.GetEmail()
		}
		result["roles"] = user.Roles
		return result
	})

	tokenService.TokenLifeTime = config.TokenLifeTime

	return tokenService, jwks
}
func generatePrivateKey(file string) *rsa.PrivateKey {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	os.WriteFile(file, privateKeyPEM, 0600)
	return privateKey
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
