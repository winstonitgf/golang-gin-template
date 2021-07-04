package main

import (
	"golang-startup/cache"
	"golang-startup/databases"
	"golang-startup/global"
	"golang-startup/router"
	"golang-startup/utils"
)

// @title Golang API
// @version 1.0
// @description Golang API 專案描述
// @termsOfService http://swagger.io/terms/

// @contact.name Winston
// @contact.email support@swagger.io

// @host localhost:8887

// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @schemes http https
func main() {

	utils.LoadEnvironment()

	databases.LoadDatabase()
	mysql, _ := global.Mysql.DB()
	defer mysql.Close()

	databases.Migrate()

	cache.LoadRedis()
	defer global.Redis.Close()

	// authorize.LoadCasbin()

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)
}
