package utils

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"strings"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB(driver string, connection string) *gorm.DB {
	if strings.EqualFold(driver, "") {
		driver = "mysql"
	}

	var err error
	sqlDB, err := sql.Open(driver, connection)
	x, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if nil != err {
		log.Error("init" + err.Error())
	}

	return x
}

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
