package cors

import (
	"bao-bet365-api/model/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.MaxAge = 7200
	config.AllowOrigins = env.Enviroment.Cors.Allow.Origins
	config.AllowHeaders = env.Enviroment.Cors.Allow.Headers
	return cors.New(config)
}