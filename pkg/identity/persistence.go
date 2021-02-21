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

func dbCheckUserExists(c *gin.Context, conn *pgx.Conn, email string) (User, error) {

	user := User{}
	query := `
		SELECT
			email, password
		FROM
			users
		WHERE
			email=$1
	`

	err := conn.QueryRow(context.Background(), query, email).Scan(
		&user.Email, &user.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "pleas try again later",
		})
	}

	return user, err
}
