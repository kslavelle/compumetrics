package tokenserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gateway/pkg/auth"
)

func Token(c *gin.Context) {

	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "please provide an email",
		})
		return
	}
	token, err := auth.GenerateToken(email)
	if err != nil {
		log.Printf("failed to create token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
