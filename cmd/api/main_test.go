package main

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/appservice"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/item"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/status"
	"testing"
)

type MockItemHandler struct {
	item.Handler
}

type MockStatusHandler struct {
	status.Handler
}

func TestCreateApp(t *testing.T) {
	itemHandler := &MockItemHandler{}
	statusHandler := &MockStatusHandler{}

	app := appservice.NewApp(itemHandler, statusHandler)

	if app == nil {
		t.Fatalf("Failed to create App, expected App, but got nil")
	}
}
