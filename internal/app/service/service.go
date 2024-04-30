package service

import (
	"context"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/db"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/handler"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
)

type App struct {
	DB db.DBTX
}

func NewApp() *App {
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

	return &App{
		DB: conn,
	}
}

func (a *App) Start() {

	handler := &handlers.Handler{
		DB: a.DB,
	}
	r := NewRouter(handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
