package identity

import "github.com/gin-gonic/gin"

func CreateIdentityProvider() *gin.Engine {
	r := gin.Default()
	r.Use(addDatabaseConnection())

	r.GET("/health", healthCheck)
	r.POST("/register", register)

	return r
}
