package main

import (
	"testing"
)

// Suppressing GoLand's nil dereference warning.
// The createApp function always returns a non-nil value. If not, the test should fail.
//
//goland:noinspection GoDfaNilDereference
func TestCreateApp(t *testing.T) {
	app := createApp()

	if app == nil {
		t.Errorf("Expected app to be not nil")
	} else {
		t.Logf("App is not nil")
	}

	if app.Handler == nil {
		t.Errorf("Expected app.Handler to be not nil")
	} else {
		t.Logf("Handler is not nil")
	}

	if app.Server == nil {
		t.Errorf("Expected app.Server to be not nil")
	} else {
		t.Logf("Server is not nil")
	}
}
