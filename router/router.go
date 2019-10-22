package route

import (
	"io"
	"os"

	. "template/main/apis"
	"template/main/middleware/token"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRoute() *gin.Engine {

	// 紀錄Log
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.SetMode(viper.GetString("mode"))
	router := gin.Default()

	// 跨網域設定
	config := cors.DefaultConfig()
	config.AllowOrigins = viper.GetStringSlice("cors.allow.origins")
	config.AllowHeaders = viper.GetStringSlice("cors.allow.headers")
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	apiv1 := router.Group("/api/v1")

	apiv1.Use(token.TokenAuthMiddleware())
	{
		apiv1.GET("/template", TemplateApi)
	}

	return router
}
