package identity

import "github.com/gin-gonic/gin"

func CreateIdentityProvider() *gin.Engine {
	r := gin.Default()
	r.GET("/health", healthCheck)

	return r
}
