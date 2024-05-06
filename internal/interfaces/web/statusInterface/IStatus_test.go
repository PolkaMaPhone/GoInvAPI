package statusInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoutes(t *testing.T) {
	testCases := []struct {
		name           string
		method         string
		route          string
		expectedStatus int
		expectedBody   string
	}{
		{name: "HandleStatus", method: http.MethodGet, route: "/status", expectedStatus: http.StatusOK, expectedBody: "Server is up and running"},
		{name: "NonExistentRoute", method: http.MethodGet, route: "/api/non_existent_route", expectedStatus: http.StatusNotFound, expectedBody: "404 page not found\n"},
		{name: "MethodNotAllowed", method: http.MethodPost, route: "/status", expectedStatus: http.StatusMethodNotAllowed, expectedBody: "method 'POST' is not allowed for route '/status'\n"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.route, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			router := chi.NewRouter()

			h := NewStatusHandler()

			r := &customRouter.CustomRouter{
				Mux: router,
			}

			h.HandleRoutes(r)

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
