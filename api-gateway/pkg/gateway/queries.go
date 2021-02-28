package gateway

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func dbConn(c *gin.Context) (*pgxpool.Pool, bool) {
	conn, ok := c.MustGet("conn").(*pgxpool.Pool)
	return conn, ok
}

func getApplication(c *pgxpool.Pool, app string, addr *string, port *int) error {

	query := `
		SELECT
			container_name, container_port
		FROM
			applications
		WHERE
			application_name=$1
	`
	return c.QueryRow(context.Background(), query, app).Scan(addr, port)
}
