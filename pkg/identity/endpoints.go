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

func loginEndpoint(c *gin.Context) {

	conn, ok := dbConn(c)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "Internal service error",
		})
		return
	}

	username := c.PostForm("username")
	hashedPassword := hashPassword(c.PostForm("password"))

	// See if user is in the database.
	user, err := getUser(c, conn, username)
	if err != nil {
		return
	}

	// Check the supplied password is correct
	if user.PasswordHash == hashedPassword {
	} // What do we do here? I know if the password is correct we need to give the user
	// a token, but I forgot how this works.

	// Return this block if the return int he block above isn't hit.
	//	return c.JSON(http.StatusUnauthorized, gin.H{"detail": "User or password not recognised."})
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
	// dbUser
	exists, err := UserInDB(c, conn, user.Email)
	// Slightly confused here. The only type of error I can find in the documentation is errornorows.
	// I guess other errors would be mistakes in the implied schema? But unsure of this as Scan apparently
	// doesn't return any other error types?
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			break
		default:
			return
		}
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{
			"detail": "There is already an account with that email."})
	}

	// // perform some checks on the email & password
	// // []
	eightOrMore, containsNumber, containsUpper, containsSpecialChar := verifyPassword(user.Password)

	if !eightOrMore {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "password must be eight or more characters.",
		})
		return
	}

	if !containsNumber {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "password must contain at least one number.",
		})
		return
	}

	if !containsUpper {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "password must contain at least one capital letter.",
		})
		return
	}

	if !containsSpecialChar {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "password must contain at least one special character.",
		})
		return
	}

	user.Password = hashPassword(user.Password)

	// insert the user into the database
	// []
	err = DBAddUser(c, conn, user)
	if err != nil {
		return
	}

	// set the context if successful
	c.JSON(http.StatusOK, gin.H{
		"detail": "account created",
	})
}
