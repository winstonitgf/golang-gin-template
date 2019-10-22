package token

import (
	"template/main/models/resp"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func respondWithError(code int, message string, c *gin.Context) {

	var respModel resp.RespModel
	respModel.IsSuccess = false
	respModel.Error = message

	c.JSON(code, respModel)

	c.Abort()
}
