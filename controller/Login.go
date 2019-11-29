package Login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index (context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}

