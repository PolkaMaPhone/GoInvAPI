package appservice

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/interfaces/web/status"
	"net/http"
	"testing"
	"time"
)

func TestAppStart(t *testing.T) {

	handler := status.NewStatusHandler()
	app := NewApp(handler)

	app.Start()

	// We need to give the server a bit of time to start up
	time.Sleep(time.Second)

	// Now we can send a request to the server
	resp, err := http.Get("http://localhost:8080/status")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}

	// Check that we get a 200 OK response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", resp.StatusCode)
	}

	// You can add more checks here to verify the response body, headers, etc.

	// Shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.Stop(ctx)
}
