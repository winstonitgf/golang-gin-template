package apis

import (
	"github.com/gin-gonic/gin"
)

func TemplateApi(c *gin.Context) {

	c.JSON(200, "Works")
	return
}
