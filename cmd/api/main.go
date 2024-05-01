package main

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/appservice"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/item"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/status"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func createApp() *appservice.App {
	config, err := dbconn.LoadConfigFile()
	if err != nil {
		log.Fatalf("Unable to load config: %v\n", err)
	}

	db := &dbconn.PgxDB{}
	_, err = dbconn.New(config, db)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// declare multiple handlers here
	itemHandler := item.NewItemHandler(db.Pool)
	statusHandler := status.NewStatusHandler()
	app := appservice.NewApp(itemHandler, statusHandler)

	return app
}

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
