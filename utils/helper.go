package utils

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/labstack/gommon/log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(driver string, connection string) *gorm.DB {
	if strings.EqualFold(driver, "") {
		driver = "mysql"
	}
	var err error
	var x *gorm.DB

	sqlDB, err := sql.Open(driver, connection)

	switch driver {
	case "mysql":

		x, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

		break
	case "pgx":
		x, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

		break

	}

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

func HashAccessToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	buf := h.Sum(nil)
	buf = buf[:len(buf)/2]
	code := base64.URLEncoding.EncodeToString(buf)
	return strings.TrimRight(code, "=")
}

func GeneratID() snowflake.ID {
	node, _ := snowflake.NewNode(1)
	return node.Generate()
}

func GeneratIDString() string {
	return GeneratID().String()
}

// GenerateRandomString generates a random string of a given length consisting of letters and digits.
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+=-"
	result := make([]byte, length)
	//rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
