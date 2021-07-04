package utils

import (
	"fmt"
	"golang-startup/global"
	"golang-startup/structs"
	"strings"

	"github.com/spf13/viper"
)

func LoadEnvironment() {
	loadConfigFile()
	viperConfigToModel()
}

func loadConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("沒有發現 config.yml，改抓取環境變數")
			viper.AutomaticEnv()
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		} else {
			// 有找到 config.yml 但是發生了其他未知的錯誤
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func viperConfigToModel() {
	var config structs.EnviromentModel
	config.Database.Client = viper.GetString("database.client")
	config.Database.Host = viper.GetString("database.host")
	config.Database.User = viper.GetString("database.user")
	config.Database.Password = viper.GetString("database.password")
	config.Database.Db = viper.GetString("database.name")
	config.Database.MaxIdle = uint(viper.GetInt("database.max_idle"))
	config.Database.MaxOpenConn = uint(viper.GetInt("database.max_open_conn"))
	config.Database.MaxLifeTime = viper.GetString("database.max_life_time")
	config.Database.Params = viper.GetString("database.params")
	config.Database.Port = viper.GetString("database.port")
	config.Database.LogEnable = viper.GetInt("database.log_enable")
	config.Database.IsMigrate = viper.GetBool("database.is_migrate")

	config.Server.Mode = viper.GetString("server.mode")
	config.Server.Port = viper.GetString("server.port")

	config.Redis.Url = strings.Replace(viper.GetString("redis.url"), "redis://", "", 1)
	config.Redis.Password = viper.GetString("redis.password")
	config.Redis.Database = viper.GetInt("redis.database")
	config.Redis.Idle = viper.GetInt("redis.idle")
	config.Redis.Active = viper.GetInt("redis.active")
	config.Redis.Protocol = viper.GetString("redis.protocol")
	config.Redis.Expire = viper.GetInt("redis.expire")
	config.Redis.Unit = viper.GetInt("redis.unit")

	config.Cors.Allow.Headers = viper.GetStringSlice("cors.allow.headers")
	config.Cors.DefaultAllowUrl = viper.GetStringSlice("cors.default_allow_url")

	global.EnvConfig = &config
}
