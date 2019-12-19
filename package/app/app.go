package app

import (
	"bao-bet365-api/model/http"

	"github.com/gin-gonic/gin"
)

type Http struct {
	C *gin.Context
}

func (this *Http) Response(data interface{}) {
	this.C.JSON(200, data)
}

func (this *Http) ErrorResponse(httpCode, code int, data interface{}) {
	this.C.JSON(httpCode, &http.ResponseModel{
		MessageCode: code,
		Data:        data,
	})
	return
}
