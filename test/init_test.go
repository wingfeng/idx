package test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/wingfeng/idx/models"
	idxmodels "github.com/wingfeng/idx/models"
	idxutils "github.com/wingfeng/idx/utils"
)

func initTestDb() *gorm.DB {
	// 	dbDriver := "mysql"
	// dbConnection := "root:password1@tcp(localhost:3306)/idx?&parseTime=true"
	dbDriver := "pgx"
	dbConnection := "host=localhost user=root password=pass@word1 dbname=idx port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db := idxutils.GetDB(dbDriver, dbConnection)
	models.Sync2Db(db)
	return db
}
func init_router() *gin.Engine {

	//初始化DB
	db := initTestDb()
	idxmodels.Sync2Db(db)

	return nil

}
func TestNewID(t *testing.T) {
	id := idxutils.GeneratID()
	t.Logf("%d", id)
}
