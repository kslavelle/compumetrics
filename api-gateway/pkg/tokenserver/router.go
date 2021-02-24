package tokenserver

import (
	"github.com/gin-gonic/gin"
)

func CreateTokenServer() *gin.Engine {

	router := gin.Default()
	router.POST("/token", Token)

	return router
}
