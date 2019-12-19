package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": http.StatusMethodNotAllowed,
			"data": "",
			"message":  "StatusMethodNotAllowed",
		})
		c.Abort()
		return
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"data":  "",
			"message": "StatusNotFound",
		})
		c.Abort()
		return
	}
}