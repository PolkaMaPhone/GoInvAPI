package main

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/appservice"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/itemInterface"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/statusInterface"
	"testing"
)

type MockItemHandler struct {
	itemInterface.Handler
}

type MockStatusHandler struct {
	statusInterface.Handler
}

func TestCreateApp(t *testing.T) {
	itemHandler := &MockItemHandler{}
	statusHandler := &MockStatusHandler{}

	app := appservice.NewApp(itemHandler, statusHandler)

	if app == nil {
		t.Fatalf("Failed to create App, expected App, but got nil")
	}
}
