package middlewares

import (
	"99live-cms-golang-api/env"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsService struct{}

func (CorsService) Cors() gin.HandlerFunc {

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.MaxAge = 7200
	config.AllowOriginFunc = func(origin string) bool {

		match := false
		// 重新在抓白名單from redis
		latestOrigins := loadWhitelist()

		// 檢查是否有符合
		for _, latestOrigin := range latestOrigins {
			if latestOrigin == origin {
				match = true
			}
		}

		return match
	}
	config.AllowHeaders = env.Config.Cors.Allow.Headers
	return cors.New(config)
}

func loadWhitelist() (origins []string) {
	origins = append(origins, env.Config.Cors.DefaultAllowUrl)

	return
}
