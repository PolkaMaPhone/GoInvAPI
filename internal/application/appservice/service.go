package appservice

import (
	"context"
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Handlers []interfaces.Handler
	Server   *http.Server
}

func NewApp(handlers ...interfaces.Handler) *App {
	router := mux.NewRouter()
	for _, h := range handlers {
		h.HandleRoutes(router)
	}

	return &App{
		Handlers: handlers,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
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
