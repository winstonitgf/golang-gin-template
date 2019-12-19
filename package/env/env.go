package env

import (
	"bao-bet365-api/model/env"
	"bao-bet365-api/package/log"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	// if _, err := os.Stat("./config.yml"); os.IsNotExist(err) {
	// 	viper.AutomaticEnv()
	// 	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// } else {
	// 	viper.SetConfigName("config")
	// 	viper.AddConfigPath(".")
	// 	err := viper.ReadInConfig()
	// 	if err != nil { // Handle errors reading the config file
	// 		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	// 	}
	// }

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			// fmt.Println(ok, err.Error())
			// 讀取不到config檔，改成讀取環境變數
			viper.AutomaticEnv()
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		} else {
			// fmt.Println("Fatal error config file")
			log.Error(err.Error())
			// 讀到config檔，但有未知錯誤
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	// 成功讀取，開始監控(可在run time的時後修改config)
	if gin.DebugMode == gin.Mode() {
		viper.WatchConfig()
	}

	if err := viper.Unmarshal(&env.Enviroment); err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s \n", err))
	}

	env.Enviroment.Server.Port = viper.GetString("server.port")

	env.Enviroment.Database.User = viper.GetString("database.user")
	env.Enviroment.Database.Password = viper.GetString("database.password")
	env.Enviroment.Database.Host = viper.GetString("database.host")
	env.Enviroment.Database.Name = viper.GetString("database.name")
	env.Enviroment.Database.Port = viper.GetString("database.port")
	env.Enviroment.Database.Parameter = viper.GetString("database.parameter")
	env.Enviroment.Database.MaxOpenConns = viper.GetInt("database.maxopenconns")
	env.Enviroment.Database.MaxIdleConns = viper.GetInt("database.maxidleconns")
	env.Enviroment.Database.MaxLifetime = viper.GetInt("database.maxlifetime")

	env.Enviroment.Mode = viper.GetString("mode")

	env.Enviroment.Redis.Url = viper.GetString("redis.url")
	env.Enviroment.Redis.Password = viper.GetString("redis.password")
	env.Enviroment.Redis.Database = viper.GetInt("redis.database")

	env.Enviroment.Cors.Allow.Origins = viper.GetStringSlice("cors.allow.origins")
	env.Enviroment.Cors.Allow.Headers = viper.GetStringSlice("cors.allow.headers")

	env.Enviroment.Ag.Actype = viper.GetString("ag.actype")

	env.Enviroment.Jwt.Secret = viper.GetString("jwt.secret")
	env.Enviroment.Jwt.Refresh = viper.GetInt("jwt.refresh")
	env.Enviroment.Jwt.Expired = viper.GetInt("jwt.expired")
}
