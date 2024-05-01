package tests

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	DbSchema := os.Getenv("DB_SCHEMA")
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", DbUser, DbPassword, DbHost, DbPort, DbName, DbSchema)

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// connect and query the database
	_, err = conn.Query(context.Background(), "SELECT * FROM items")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
}
