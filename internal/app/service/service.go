package service

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/apihandler"
	"log"
	"net/http"
)

type App struct {
	Handler *apihandler.APIHandler
}

func NewApp(handler *apihandler.APIHandler) *App {
	return &App{
		Handler: handler,
	}
}

func (a *App) Start() {
	r := NewRouter(a.Handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
