package cache

import (
	"golang-startup/global"
	"time"

	"github.com/gomodule/redigo/redis"
)

func LoadRedis() {
	global.Redis = &redis.Pool{
		MaxIdle:         global.EnvConfig.Redis.Idle,
		MaxActive:       global.EnvConfig.Redis.Active,
		IdleTimeout:     1 * time.Hour,
		Wait:            true,
		MaxConnLifetime: 2 * time.Hour,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(global.EnvConfig.Redis.Protocol, global.EnvConfig.Redis.Url)
			if err != nil {
				return nil, err
			}

			if global.EnvConfig.Redis.Password != "" {
				_, err = c.Do("AUTH", global.EnvConfig.Redis.Password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", global.EnvConfig.Redis.Database); err != nil {
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
