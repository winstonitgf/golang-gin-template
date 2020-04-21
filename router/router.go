package router

import (
	"99live-cms-golang-api/controller/auth"
	"99live-cms-golang-api/env"
	"99live-cms-golang-api/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {

	// 設定Gin Mode (Local=>debug、測試機=>test、正式機=>production)
	gin.SetMode(env.Config.Mode)
	router := gin.Default()

	// 找不到方法或route的時後，統一處理
	// router.NoMethod(common.NoMethodHandler())
	// router.NoRoute(common.NoRouteHandler())

	// 跨網域設定
	var corsService middlewares.CorsService
	router.Use(corsService.Cors())

	apiv1 := router.Group("/api/v1")

	apiAuth := apiv1.Group("/auth")
	apiAuth.POST("/login", auth.Login)
	apiAuth.POST("/verify", auth.Verify)

	return router
}
