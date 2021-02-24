package tokenserver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gateway/pkg/auth"
)

func Token(c *gin.Context) {

	email := c.PostForm("email")
	token, err := auth.GenerateToken(email)
	if err != nil {
		// log
		// error check
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
