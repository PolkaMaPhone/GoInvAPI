package itemInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemIntegration(t *testing.T) {
	testCases := []struct {
		name   string
		method string
		route  string
		status int
	}{
		{name: "HandleGet", method: http.MethodGet, route: "/items/1", status: http.StatusOK},
		{name: "HandleGetAll", method: http.MethodGet, route: "/items", status: http.StatusOK},
		{name: "HandleGet_NonExistentItem", method: http.MethodGet, route: "/items/9999", status: http.StatusNotFound},
		{name: "HandleGet_InvalidID", method: http.MethodGet, route: "/items/invalid", status: http.StatusBadRequest},
		{name: "NotAllowedMethod", method: http.MethodPost, route: "/items/1", status: http.StatusMethodNotAllowed},
	}

	config, err := dbconn.LoadConfigFile()
	if err != nil {
		log.Fatalf("Unable to load config: %v\n", err)
	}

	db := &dbconn.PgxDB{}
	_, err = dbconn.GetPoolInstance(config, db)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	itemRepo := itemDomain.NewRepository(db.Pool)
	itemService := itemDomain.NewService(itemRepo)
	itemHandler := NewItemHandler(itemService)

	router := chi.NewRouter()
	r := &customRouter.CustomRouter{
		Mux: router,
	}
	itemHandler.HandleRoutes(r)

	server := httptest.NewServer(router)
	defer server.Close()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.NewRequest(tt.method, server.URL+tt.route, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			client := &http.Client{}
			response, err := client.Do(resp)
			if err != nil {
				t.Fatalf("could not send request: %v", err)
			}

			if response.StatusCode != tt.status {
				t.Errorf("expected status %d, got %d", tt.status, response.StatusCode)
			}
		})
	}
}
