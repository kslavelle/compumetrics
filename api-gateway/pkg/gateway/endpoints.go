package gateway

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"api-gateway/pkg/auth"
)

func ProxyRequest(c *gin.Context) {

	conn, ok := dbConn(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "internal server error",
		})
		return
	}

	email, err := auth.ValidateToken(c)
	if err != nil {
		return
	}

	application := c.Param("application")
	var containerName string
	var containerPort int

	err = getApplication(conn, application, &containerName, &containerPort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "internal server error",
		})
		return
	}

	// forwards the request on with the *path
	address := fmt.Sprintf("http://%v:%v", containerName, containerPort)
	applicationPart := fmt.Sprintf("/%v", application)

	trimmedPath := strings.Replace(c.Request.URL.Path, applicationPart, "", 1)
	c.Request.URL.Path = trimmedPath

	proxiedURL, _ := url.Parse(address)
	target := httputil.NewSingleHostReverseProxy(proxiedURL)

	c.Request.Header.Set("X-Authenticated-UserId", email)
	target.ServeHTTP(c.Writer, c.Request)
}
