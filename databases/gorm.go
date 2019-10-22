package databases

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var Eloquent *gorm.DB

func GormOpen() {
	var err error
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	name := viper.GetString("database.name")
	Eloquent, err = gorm.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	if Eloquent.Error != nil {
		panic(Eloquent.Error)
	}

	// debug mode 的時後，開啟log
	Eloquent.LogMode(gin.DebugMode == gin.Mode())
}

func GormClose() {
	Eloquent.Close()
}

// ALTER USER 'root'@'%' IDENTIFIED BY '1qaz2wsx' PASSWORD EXPIRE NEVER;

// ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '1qaz2wsx';

// FLUSH PRIVILEGES;
