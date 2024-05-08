package locationInterface

import (
	"errors"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/locationDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoutes(t *testing.T) {
	testCases := []struct {
		name          string
		method        string
		route         string
		locationID    int
		err           error
		expStatus     int
		mockSetupFunc func(ms *MockService)
	}{
		{name: "HandleGet", method: http.MethodGet, route: "/api/locations/%d", locationID: 1, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetLocationByID", int32(1)).Return(&locationDomain.Location{LocationID: 1}, nil)
		}},
		{name: "HandleGet_Error", method: http.MethodGet, route: "/api/locations/%d", locationID: 1, err: errors.New("some error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetLocationByID", int32(1)).Return(&locationDomain.Location{}, errors.New("some error"))
		}},
		{name: "HandleGet_InvalidID", method: http.MethodGet, route: "/api/locations/invalid", locationID: 0, err: nil, expStatus: http.StatusBadRequest, mockSetupFunc: func(ms *MockService) {}},
		{name: "HandleGetAll_Error", method: http.MethodGet, route: "/api/locations", locationID: 0, err: errors.New("some error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetAllLocations").Return([]*locationDomain.Location{}, errors.New("some error"))
		}},
		//{name: "HandleGet_NotFound", method: http.MethodGet, route: "/locations/%d", locationID: 999, err: errors.New("the parameter 'location' with id '999' returned no results"), expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {
		//	ms.On("GetLocationByID", int32(999)).Return(nil, errors.New("the parameter 'location' with id '999' returned no results"))
		//}},
		//{name: "HandleGetAll_NoLocations", method: http.MethodGet, route: "/locations", locationID: 0, err: errors.New("no locations found"), expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {
		//	ms.On("GetAllLocations").Return(nil, errors.New("no locations found"))
		//}},
		{name: "NonExistentRoute", method: http.MethodGet, route: "/api/non_existent_route", locationID: 0, err: nil, expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {}},
		{name: "NotAllowedMethod", method: http.MethodPost, route: "/api/locations/%d", locationID: 1, err: &utils.MethodNotAllowedError{Method: "POST", Route: "/api/locations/%d"}, expStatus: http.StatusMethodNotAllowed, mockSetupFunc: func(ms *MockService) {}},
		{name: "ServiceError", method: http.MethodGet, route: "/api/locations/%d", locationID: 1, err: errors.New("internal server error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetLocationByID", int32(1)).Return(nil, errors.New("some other error"))
		}},
		{name: "NoResults", method: http.MethodGet, route: "/api/locations/%d", locationID: 999, err: &utils.NoResultsForParameterError{ParameterName: "location_id", ID: "999"}, expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {
			ms.On("GetLocationByID", int32(999)).Return(nil, pgx.ErrNoRows)
		}},
		{name: "UnexpectedError", method: http.MethodGet, route: "/api/locations/%d", locationID: 1, err: errors.New("unexpected error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetLocationByID", int32(1)).Return(nil, errors.New("unexpected error"))
		}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			mockService := &MockService{}
			service := locationDomain.NewService(mockService)
			handler := NewLocationHandler(service)
			tt.mockSetupFunc(mockService)

			route := tt.route
			if tt.locationID != 0 {
				route = fmt.Sprintf(tt.route, tt.locationID)
			}

			req, err := http.NewRequest(tt.method, route, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			responseRecorder := httptest.NewRecorder()
			router := chi.NewRouter()

			r := &customRouter.CustomRouter{
				Mux: router,
			}
			handler.HandleRoutes(r)
			router.ServeHTTP(responseRecorder, req)

			// Log request
			log.Printf("Request Method: %s\n", req.Method)
			log.Printf("Request URL: %s\n", req.URL)
			log.Printf("Request Headers: %s\n", req.Header)

			// Log response
			log.Printf("Response Status Code: %d\n", responseRecorder.Code)
			log.Printf("Response Body: %s\n", responseRecorder.Body.String())

			// Check the status code
			if status := responseRecorder.Code; status != tt.expStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expStatus)
			}

			// Check the error returned by the handler
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("handler returned wrong error: got %v want %v", err, tt.err)
			}
		})
	}
}
