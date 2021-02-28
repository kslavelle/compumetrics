package gateway

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func connectToDatabase() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("error connecting to DB: %v\n", err)
	}
	return pool
}

func addDatabaseConnection() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := connectToDatabase()
		c.Set("conn", conn)
		c.Next()
	}
}
