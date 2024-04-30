// Description: This file is the entry point of the application.
// It creates a new instance of the service and starts it with the HandleRequest function from the handlers package.
package main

import (
	handlers "github.com/PolkaMaPhone/GoInvAPI/internal/app/handler"
	services "github.com/PolkaMaPhone/GoInvAPI/internal/app/service"
)

func main() {
	app := services.NewApp()
	app.Start(handlers.HandleRequest)
}
