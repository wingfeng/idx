package test

import (
	"github.com/gin-gonic/gin"

	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/utils"
)

func init_router() *gin.Engine {
	//"pgx", "host=localhost user=postgres password=pass@word1 dbname=idx port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	// dbDriver := "mysql"
	// dbConnection := "root:eATq1GDhsP@tcp(localhost:31332)/idx?&parseTime=true"
	dbDriver := "pgx"
	dbConnection := "host=localhost user=root password=pass@word1 dbname=idx port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// hashKey := []byte("FF51A553-72FC-478B-9AEF-93D6F506DE91")
	// session.InitManager(
	// 	session.SetStore(
	// 		cookie.NewCookieStore(
	// 			cookie.SetCookieName("demo_cookie_store_id"),
	// 			cookie.SetHashKey(hashKey),
	// 		),
	// 	),

	//初始化DB
	db := utils.GetDB(dbDriver, dbConnection)
	idxmodels.Sync2Db(db)

	return nil

}
