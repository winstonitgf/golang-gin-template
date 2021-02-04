package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 登入
// @Tags OAUTH
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 400 {object} models.User
// @Security ApiKeyAuth
// @Router /auth/login [post]
func Login(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}
