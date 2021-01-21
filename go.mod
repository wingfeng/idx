module github.com/wingfeng/idx

go 1.15

replace github.com/go-oauth2/oauth2/v4 v4.2.0 => /home/wing/git_repos/oauth2

require (
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-oauth2/oauth2/v4 v4.2.0
	github.com/go-session/cookie v3.0.1+incompatible // indirect
	github.com/go-session/gorm v0.0.0-20190223144354-7d5f87dcd6c3
	github.com/go-session/session v3.1.2+incompatible
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/labstack/gommon v0.3.0
	github.com/magiconair/properties v1.8.1
	github.com/rs/cors v1.7.0
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/guregu/null.v4 v4.0.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.9
)
