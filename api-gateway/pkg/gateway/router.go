package gateway

import (
	"github.com/gin-gonic/gin"
)

func CreateGatewayServer() *gin.Engine {

	router := gin.Default()
	router.Use(addDatabaseConnection())

	router.POST("/:application/*path", ProxyRequest)

	return router
}
