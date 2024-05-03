package main

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/appservice"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/categoryDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/groupDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/locationDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/categoryInterface"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/groupInterface"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/itemInterface"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/locationInterface"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/statusInterface"
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
	_, err = dbconn.GetPoolInstance(config, db)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Create an instance of the repositories
	itemRepo := itemDomain.NewRepository(db.Pool)
	categoryRepo := categoryDomain.NewRepository(db.Pool)
	groupRepo := groupDomain.NewRepository(db.Pool)
	locationRepo := locationDomain.NewRepository(db.Pool)

	// Create services from repositories
	itemService := itemDomain.NewService(itemRepo)
	categoryService := categoryDomain.NewService(categoryRepo)
	groupService := groupDomain.NewService(groupRepo)
	locationService := locationDomain.NewService(locationRepo)

	// Create handlers from services
	itemHandler := itemInterface.NewItemHandler(itemService)
	categoryHandler := categoryInterface.NewCategoryHandler(categoryService)
	groupHandler := groupInterface.NewGroupHandler(groupService)
	locationHandler := locationInterface.NewLocationHandler(locationService)

	// Create status handler
	statusHandler := statusInterface.NewStatusHandler()

	// Create GetPoolInstance App and inject handlers
	app := appservice.NewApp(
		itemHandler,
		categoryHandler,
		statusHandler,
		groupHandler,
		locationHandler)

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
