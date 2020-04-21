package auth

import (
	"99live-cms-golang-api/models"
	"99live-cms-golang-api/packages"
	"99live-cms-golang-api/services/auth"
	"99live-cms-golang-api/structs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var userInfo models.User
	err := c.Bind(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ip := c.ClientIP()

	var loginData structs.LoginModel
	var authService auth.AuthService
	if loginData = authService.Login(userInfo, ip); len(authService.Errors()) != 0 {

		packages.HttpErrorResponse(c, authService.Errors())
		return
	} else {

		c.JSON(http.StatusOK, loginData)
		return
	}
}

func Verify(c *gin.Context) {

	authorization := c.Request.Header.Get("Authorization")
	token := strings.Split(authorization, " ")[1]

	var authService auth.AuthService
	if userData := authService.Verify(token); len(authService.Errors()) != 0 {

		packages.HttpErrorResponse(c, authService.Errors())
		return
	} else {

		c.JSON(http.StatusOK, userData)
		return
	}
}
