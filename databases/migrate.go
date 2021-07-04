package databases

import (
	"golang-startup/global"
	"golang-startup/models"
)

func Migrate() {
	if global.EnvConfig.Database.IsMigrate {
		global.Mysql.AutoMigrate(&models.User{})
		global.Mysql.Create(&models.User{Account: "winston", Nickname: "Winston", Email: "winston.lu@bimap.co", Phone: "0912345678"})
	}
}
