package instances

import (
	"99live-cms-golang-api/env"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Mysql *gorm.DB
)

func InitDatabasePool() {
	createPool()
}

func createPool() {

	// 取得config參數
	client := env.Config.Database.Client
	host := env.Config.Database.Host
	port := env.Config.Database.Port
	user := env.Config.Database.User
	password := env.Config.Database.Password
	dbname := env.Config.Database.Db
	parameter := env.Config.Database.Params

	var err error
	Mysql, err = gorm.Open(client, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbname, parameter))
	if err != nil {
		panic(err)
	}
	// defer Mysql.Close()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	Mysql.DB().SetMaxIdleConns(int(env.Config.Database.MaxIdle))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	Mysql.DB().SetMaxOpenConns(int(env.Config.Database.MaxOpenConn))

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	lifeTime, _ := time.ParseDuration(env.Config.Database.MaxLifeTime)
	Mysql.DB().SetConnMaxLifetime(lifeTime)

	Mysql.LogMode(gin.DebugMode == gin.Mode())

}
