package instances

import (
	"99live-cms-golang-api/env"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func InitRedisPool() {
	RedisPool = createRedisPool()

}

func createRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   env.Config.Redis.Idle,
		MaxActive: env.Config.Redis.Active,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(env.Config.Redis.Protocol, env.Config.Redis.Url)
			if err != nil {
				return nil, err
			}

			if env.Config.Redis.Password != "" {
				_, err = c.Do("AUTH", env.Config.Redis.Password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", env.Config.Redis.Database); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
