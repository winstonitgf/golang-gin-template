package token

import (
	"bao-bet365-api/enum/message"
	"bao-bet365-api/model/http"
	tokenPackage "bao-bet365-api/package/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		// 找不到token 報錯
		if authorization == "" {
			respondWithError(message.UNAUTHORIZED, message.TOKEN_NOT_FOUND, c)
			return
		}
		// 抓取token
		token := strings.Split(authorization, " ")[1]
		// 驗證token
		var tokenService tokenPackage.TokenService
		err := tokenService.ValidateToken(token)
		if err != nil {
			respondWithError(message.UNAUTHORIZED, message.TOKEN_INVALID, c)
			return
		}
		c.Next()
	}
}

func respondWithError(code int, message int, c *gin.Context) {

	var respModel http.ResponseModel
	respModel.MessageCode = message

	c.JSON(code, respModel)
	c.Abort()
}
