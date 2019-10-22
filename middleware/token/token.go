package token

import (
	"template/main/models/resp"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if isDebug {
		// 	c.Next()
		// 	return
		// }
		// // 檢查有沒有token
		// apiToken := c.Request.Header.Get("api_token")
		// if apiToken == "" {

		// 	// 如果沒有Token，就檢查有沒有 refresh_token
		// 	refreshToken := c.Request.Header.Get("refresh_token")
		// 	if refreshToken == "" {
		// 		respondWithError(401, "找不到合法token", c)
		// 		return
		// 	}

		// 	// 如果有找到 refresh_token，就檢查
		// 	valid, err := services.ValidateToken(refreshToken)
		// 	if err != nil {
		// 		if err.Error() == "Token is expired" {

		// 			// 過期的話，特別回傳
		// 			respondWithError(viper.GetInt("token.expirecode"), err.Error(), c)
		// 			return
		// 		} else {
		// 			respondWithError(401, err.Error(), c)
		// 			return
		// 		}
		// 	}
		// 	if !valid {
		// 		respondWithError(401, "token驗證失敗", c)
		// 		return
		// 	}
		// 	c.Next()
		// 	return
		// }

		// // 驗證Token
		// valid, err := services.ValidateToken(apiToken)
		// if err != nil {
		// 	if err.Error() == "Token is expired" {

		// 		// 過期的話，特別回傳407
		// 		respondWithError(viper.GetInt("token.expirecode"), err.Error(), c)
		// 		return
		// 	} else {
		// 		respondWithError(401, err.Error(), c)
		// 		return
		// 	}
		// }
		// if !valid {
		// 	respondWithError(401, "token驗證失敗", c)
		// 	return
		// }
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
