package service

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/handler"
	"log"
	"net/http"
)

type App struct {
	// Add fields such as database connections, config, etc.
}

func NewApp() *App {
	// Initialize your app with necessary setup, like database connection.
	return &App{}
}

func (a *App) Start() {
	http.HandleFunc("/", handlers.HandleRequest)
	http.HandleFunc("/items", handlers.HandleGetItem) // New route for GetItems
	log.Fatal(http.ListenAndServe(":8080", nil))
}
