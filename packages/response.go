package packages

import (
	"99live-cms-golang-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpErrorResponse(c *gin.Context, errList []error) {

	var httpResponse structs.HttpResponseModel
	for _, err := range errList {
		httpResponse.Errors = append(httpResponse.Errors, err.Error())
	}
	c.JSON(http.StatusBadRequest, httpResponse)
}
