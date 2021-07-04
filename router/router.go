package router

import (
	"github.com/gin-gonic/gin"

	. "golang-startup/controller"

	_ "golang-startup/docs"
	"golang-startup/global"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter() *gin.Engine {

	gin.SetMode(global.EnvConfig.Server.Mode)
	router := gin.Default()

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := router.Group("api/v1")
	{
		apiAuth := apiv1.Group("redis")
		{
			apiAuth.POST("getset/:id", RedisGetSET)
			apiAuth.POST("hmgetset/:id", RedisHMGetSET)
			apiAuth.POST("listgetset", RedisListGetSet)
			apiAuth.POST("setgetset", RedisSetGetSet)
			apiAuth.POST("sortedsetgetset", RedisSortedSetGetSet)
		}
	}

	return router
}
