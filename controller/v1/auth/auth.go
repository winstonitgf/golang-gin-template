package auth

import (
	"bao-bet365-api/enum/message"
	"bao-bet365-api/package/app"
	"bao-bet365-api/service/auth"

	"github.com/gin-gonic/gin"
)

// @Summary 登入
// @Tags Auth
// @Description 登入
// @Accept  json
// @Produce  json
// @Param body body member.LoginRequestModel true "登入資料"
// @Success 200  {string} string "成功"
// @Failure 400  {string} string "失敗"
// @Router /auth/login [post]
func Login(c *gin.Context) {

	var authService auth.AuthService
	var http app.Http
	http.C = c

	// 登入驗證
	err := authService.Login()
	if err != nil {
		http.ErrorResponse(message.ERROR, 1004, nil)
		return
	}

	http.Response(nil)
	return
}
