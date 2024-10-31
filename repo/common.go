package repo

import (
	"github.com/wingfeng/idx/conf"
	"github.com/wingfeng/idx/utils"
	"gorm.io/gorm"
)

func DefaultDB() *gorm.DB {
	conn := conf.Default.Connection
	driver := conf.Default.Driver
	return utils.GetDB(driver, conn)
}
