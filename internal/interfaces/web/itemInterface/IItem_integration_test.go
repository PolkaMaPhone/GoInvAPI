package itemInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/go-chi/chi/v5"
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
		{name: "NotAllowedMethod", method: http.MethodPost, route: "/items/1", status: http.StatusMethodNotAllowed},
		// TODO - Add more test cases...
	}

	mockService := &MockService{}
	mockService.On("GetItemByID", int32(1)).Return(&itemDomain.Item{}, nil)
	service := itemDomain.NewService(mockService)
	handler := NewItemHandler(service)

	router := chi.NewRouter()
	r := &customRouter.CustomRouter{
		Mux: router,
	}
	handler.HandleRoutes(r)

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
