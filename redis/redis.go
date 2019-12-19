package redis

import (
	envModel "bao-bet365-api/model/env"
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

var Cache *redis.Client

func Init() {

	Cache = redis.NewClient(&redis.Options{
		Addr:     strings.Replace(envModel.Enviroment.Redis.Url, "redis://", "", 1),
		Password: envModel.Enviroment.Redis.Password, // no password set
		DB:       envModel.Enviroment.Redis.Database, // use default DB
	})
	_, err := Cache.Ping().Result()
	if err != nil {
		fmt.Println("Redis連線失敗")
	} else {
		fmt.Println("Redis連線成功")
	}

}

func Close() {
	_ = Cache.Close()
}
