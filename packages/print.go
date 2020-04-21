package packages

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Print(a ...interface{}) {
	if gin.DebugMode == gin.Mode() {
		fmt.Println(a...)
	}
}
