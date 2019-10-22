// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost:8001
//      Version: 0.0.1
//      Schemes: http
//
// swagger:meta
package main

import (
	"fmt"
	"os"
	"strings"
	gorm "template/main/databases"
	"template/main/redis"
	route "template/main/router"

	"github.com/spf13/viper"
)

func LoadConfig() {
	if _, err := os.Stat("./config.yml"); os.IsNotExist(err) {
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	fmt.Print(viper.GetString("server.port"))
}

func main() {

	// 讀取設定檔
	LoadConfig()

	// 開啟DB連線
	gorm.GormOpen()
	defer gorm.GormClose()

	// 開啟Redis連線
	redis.Open()
	defer redis.Client.Close()

	// 啟動Gin
	app := route.InitRoute()
	app.Run(viper.GetString("server.port"))
}
