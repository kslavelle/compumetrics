package identity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"detail": "ok",
	})
}

func register(c *gin.Context) {

	conn, ok := dbConn(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
		return
	}

	user := RegisterUser{}

	// email := c.Request.Form.Get("email")
	// password := c.Request.Form.Get("password")
	// user := RegisterUser{email, password}

	// is equivelent to the above
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Expected [email, password] as form data.",
		})
		return
	}

	// check that the user is not in the database
	dbUser, err := dbCheckUserExists(c, conn, user.Email)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			break
		default:
			return
		}
	}

	// perform some checks on the email & password
	// []

	user.Password = hashPassword(user.Password)

	// insert the user into the database
	// []

	// set the context if successful
	c.JSON(http.StatusOK, gin.H{
		"detail": "account created",
	})
}
