package main

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/apihandler"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := createApp()

	go app.Start()

	// Wait for an interrupt signal
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	// Shutdown the server when the interrupt signal is received
	app.Stop(context.Background())
}

func createApp() *service.App {
	handler := apihandler.NewAPIHandler()
	app := service.NewApp(handler)
	return app
}
