package identity

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func connectToDatabase() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func addDatabaseConnection() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := connectToDatabase()
		c.Set("conn", conn)
		c.Next()
	}
}
