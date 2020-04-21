package env

import (
	"99live-cms-golang-api/structs"

	"github.com/spf13/viper"
)

var Config *structs.EnviromentModel

type ViperPluginService struct {
}

func (ViperPluginService) ConfigToModel() {

	var env structs.EnviromentModel

	// 環境
	env.Server.Port = viper.GetString("server.port")
	env.Mode = viper.GetString("mode")

	// redis
	env.Redis.Url = viper.GetString("redis.url")
	env.Redis.Password = viper.GetString("redis.password")
	env.Redis.Database = viper.GetInt("redis.database")
	env.Redis.Idle = viper.GetInt("redis.idle")
	env.Redis.Active = viper.GetInt("redis.active")
	env.Redis.Protocol = viper.GetString("redis.protocol")

	// 跨網域
	env.Cors.Allow.Headers = viper.GetStringSlice("cors.allow.headers")
	env.Cors.DefaultAllowUrl = viper.GetString("cors.default_allow_url")

	// jwt
	env.Jwt.Secret = viper.GetString("jwt.secret")

	// database
	env.Database.Client = viper.GetString("database.client")
	env.Database.Host = viper.GetString("database.host")
	env.Database.User = viper.GetString("database.user")
	env.Database.Password = viper.GetString("database.passord")
	env.Database.Db = viper.GetString("database.db")
	env.Database.MaxIdle = uint(viper.GetInt("database.max_idle"))
	env.Database.MaxOpenConn = uint(viper.GetInt("database.max_life_time"))
	env.Database.MaxLifeTime = viper.GetString("database.max_open_conn")
	env.Database.Params = viper.GetString("database.params")
	env.Database.Port = viper.GetString("database.port")
	Config = &env
}
