package main

import service "github.com/PolkaMaPhone/GoInvAPI/internal/app/service"

func main() {
	app := service.NewApp()
	app.Start()
}
