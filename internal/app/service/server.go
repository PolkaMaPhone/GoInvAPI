package app

import (
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
	// Setup your HTTP server, routes, etc.
	http.HandleFunc("/", a.handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *App) handleRequest(w http.ResponseWriter, r *http.Request) {
	// Handle your requests here
	w.Write([]byte("Hello, World!"))
}
