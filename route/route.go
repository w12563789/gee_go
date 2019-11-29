package route

import (
	"geeoffice/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/test", Index.Test)
	router.GET("/test2", Login.Index)
	return router
}
