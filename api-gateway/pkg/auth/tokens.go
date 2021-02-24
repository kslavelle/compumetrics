package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(email string) (string, error) {

	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	encryptionKey := os.Getenv("ACCESS_SECRET")
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := at.SignedString([]byte(encryptionKey))
	token := fmt.Sprintf("Bearer %v", accessToken)

	return token, err
}

func ValidateToken(c *gin.Context, handler func(*gin.Context, string)) {

	bearerToken := c.GetHeader("Authorization")
	tokenString := strings.Split(bearerToken, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Undexpected signing method")
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"detail": "error validating token",
		})
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"detail": "error validating token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"detail": "error validating token",
		})
		return
	}

	email := claims["email"].(string)
	handler(c, email)
}
