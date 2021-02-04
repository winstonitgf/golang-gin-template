package main

import (
	"fmt"
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

	// 初始化 env
	utils.LoadEnvironment()
	fmt.Println("參數初始化成功...")

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)

	// cache.InitRedisPool()
	// defer cache.RedisPool.Close()

	// // 初始化 db
	// database.InitDatabasePool()
	// // database.Migration()
	// sqlDB, _ := database.Mysql.DB()
	// defer sqlDB.Close()
	// fmt.Println("資料庫始化成功...")

	// // casbin init
	// var casbinService casbinLib.CasbinService
	// casbinService.Init()

	// // 初始化 db
	// migration.InitTables()

	// r := router.Router()
	// r.Run(":3000")
}
