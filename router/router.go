package router

import (
	"bao-bet365-api/middleware/token"
	"bao-bet365-api/middleware/user"
	"io"
	"os"

	"bao-bet365-api/controller/v1/auth"
	_ "bao-bet365-api/docs"
	"bao-bet365-api/middleware/common"
	"bao-bet365-api/middleware/cors"
	"bao-bet365-api/model/env"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	// 紀錄Log
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 設定Gin Mode (Local=>debug、測試機=>test、正式機=>production)
	gin.SetMode(env.Enviroment.Mode)
	router := gin.Default()

	// 找不到方法或route的時後，統一處理
	router.NoMethod(common.NoMethodHandler())
	router.NoRoute(common.NoRouteHandler())

	// 跨網域設定
	router.Use(cors.Cors())

	// 提供container健康檢測用
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	// swagger 文件
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := router.Group("/api/v1")
	apiv1.POST("/auth/login", auth.Login)

	apiv1.Use(token.TokenAuthMiddleware(), user.UserAuthMiddleware())
	{

	}

	return router
}
