package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Client *redis.Client

func Open() {
	// redis 公版提供的redis格式，要另外處理
	addr := viper.GetString("redis.url")

	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString("redis.password"), // no password set
		DB:       0,                                 // use default DB
	})
}
