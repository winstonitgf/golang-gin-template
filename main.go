package main

import (
	"golang-startup/authorize"
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

// @host staging-api.99live.live

// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @schemes https
func main() {

	utils.LoadEnvironment()

	databases.LoadDatabase()
	mysql, _ := global.Mysql.DB()
	defer mysql.Close()

	cache.LoadRedis()
	defer global.Redis.Close()

	authorize.LoadCasbin()

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)
}
