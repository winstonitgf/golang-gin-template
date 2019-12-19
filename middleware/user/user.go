package user

import (
	"bao-bet365-api/enum/message"
	"bao-bet365-api/model/http"

	"github.com/gin-gonic/gin"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")

		// 找不到token 報錯
		if authorization == "" {
			respondWithError(message.UNAUTHORIZED, message.TOKEN_NOT_FOUND, c)
			return
		}

		// var authService auth.AuthService
		// // 取得token用戶
		// userEntity, errorModel := authService.GetUser(authorization)
		// if errorModel.Error != nil {
		// 	respondWithError(message.UNAUTHORIZED, errorModel.ErrorCode, c)
		// 	return
		// }

		c.Next()
	}
}

func respondWithError(code int, message int, c *gin.Context) {

	var respModel http.ResponseModel
	respModel.MessageCode = message

	c.JSON(code, respModel)
	c.Abort()
}
