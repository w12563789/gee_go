package Index

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Test (context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
