module github.com/wingfeng/idx

go 1.15

// replace oauth2 => ./oauth2/

require (
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	//oauth2 v4.2.0
	github.com/go-session/gorm v0.0.0-20190223144354-7d5f87dcd6c3
	github.com/go-session/session v3.1.2+incompatible
	github.com/google/uuid v1.1.5
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/labstack/gommon v0.3.0
	github.com/magiconair/properties v1.8.1
	github.com/rs/cors v1.7.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.1
	github.com/tidwall/buntdb v1.1.7
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	golang.org/x/text v0.3.3 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/guregu/null.v4 v4.0.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.9
)
