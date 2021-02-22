package identity

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func dbConn(c *gin.Context) (*pgx.Conn, bool) {
	conn, ok := c.MustGet("conn").(*pgx.Conn)
	return conn, ok
}

func getUser(c *gin.Context, conn *pgx.Conn, username string) (User, error) {

	user := User{}
	query := `
		SELECT
			email, hashed_password
		FROM
			users
		WHERE
			username=$1
	`
	err := conn.QueryRow(
		context.Background(), query, username,
	).Scan(&user.Email, &user.PasswordHash)

	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"detail": "invalid username or password",
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"detail": "Internal service error.",
			})
		}
		return user, err
	}
	return user, nil
}

// UserInDB Checks if the User is in the database.
func UserInDB(c *gin.Context, conn *pgx.Conn, email string) (exists bool, err error) {

	query := `
        SELECT
            count(*) > 0
        FROM
            users
        WHERE
            email=$1
    `
	err = conn.QueryRow(
		context.Background(), query, email,
	).Scan(&exists)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
	}
	return
}

// DBAddUser Add email and hashed password to SQL DB.
func DBAddUser(c *gin.Context, conn *pgx.Conn, user RegisterUser) error {
	query := `
		INSERT INTO users
			(email, hashed_password)
		VALUES
			($1, $2)
	`
	_, err := conn.Exec(context.Background(), query, user.Email, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": "unknown",
		})
		return err
	}

	return nil
}
