package main

import (
	services "github.com/PolkaMaPhone/GoInvAPI/internal/app/service"
)

func main() {
	app := services.NewApp()
	app.Start()
}
