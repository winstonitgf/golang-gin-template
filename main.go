package main

import (
	"fmt"
	"strings"

	"99live-cms-golang-api/env"
	"99live-cms-golang-api/packages"

	"99live-cms-golang-api/router"

	"99live-cms-golang-api/plugins/instances"

	"github.com/spf13/viper"
)

func main() {

	loadConfig()

	instances.InitRedisPool()
	packages.Print("Redis: " + env.Config.Redis.Url + " 連線池初始化完成...")

	instances.InitDatabasePool()
	packages.Print("Mysql: " + env.Config.Database.Host + ":" + env.Config.Database.Port + " 連線池初始化完成...")

	app := router.InitRoute()
	packages.Print("路由器初始化完成....")

	app.Run(env.Config.Server.Port)
}

func loadConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			// 進到這邊代表找不到 config.yml
			// 找不到 config.yml 的話就抓取環境變數
			viper.AutomaticEnv()
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		} else {

			// 有找到 config.yml 但是發生了其他未知的錯誤
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	} else {
		var viperPluginService env.ViperPluginService
		viperPluginService.ConfigToModel()
	}

	packages.Print("讀取 config.yml 完成")

	return
}
