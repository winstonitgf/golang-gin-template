// @title 保哥-前台-API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http https
// @host 31-review-develop-82qm1o.atavral.ninja
// @BasePath /api/v1
package main

import (
	"bao-bet365-api/database"
	"bao-bet365-api/model/env"
	"bao-bet365-api/package/app"
	_ "bao-bet365-api/package/env"
	"bao-bet365-api/redis"

	"bao-bet365-api/router"
)

func main() {

	var appInit app.AppInit
	appInit.Init()

	defer redis.Close()
	defer database.Close()

	// 啟動Gin
	app := router.InitRouter()
	app.Run(env.Enviroment.Server.Port)
}
