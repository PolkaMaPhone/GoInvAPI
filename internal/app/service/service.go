package service

import (
	"context"
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/apihandler"
	"log"
	"net/http"
)

type App struct {
	Handler *apihandler.APIHandler
	Server  *http.Server
}

func NewApp(handler *apihandler.APIHandler) *App {
	return &App{
		Handler: handler,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: NewRouter(handler),
		},
	}
}

func (a *App) Start() {
	go func() {
		if err := a.Server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			// Unexpected server shutdown
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
}

func (a *App) Stop(ctx context.Context) {
	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
