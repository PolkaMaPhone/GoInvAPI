package main

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/apihandler"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/service"
)

func main() {
	apiHandler := apihandler.NewAPIHandler()
	app := service.NewApp(apiHandler)
	app.Start()
}
