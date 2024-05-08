package groupInterface

import (
	"errors"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/groupDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetGroupByID(id int32) (*groupDomain.Group, error) {
	args := m.Called(id)
	group, ok := args.Get(0).(*groupDomain.Group)
	if !ok {
		return nil, args.Error(1)
	}
	return group, args.Error(1)
}

func (m *MockService) GetAllGroups() ([]*groupDomain.Group, error) {
	args := m.Called()
	return args.Get(0).([]*groupDomain.Group), args.Error(1)
}

func TestHandleRoutes(t *testing.T) {
	testCases := []struct {
		name          string
		method        string
		route         string
		groupID       int
		err           error
		expStatus     int
		mockSetupFunc func(ms *MockService)
	}{
		{name: "HandleGet", method: http.MethodGet, route: "/api/groups/%d", groupID: 1, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetGroupByID", int32(1)).Return(&groupDomain.Group{GroupID: 1}, nil)
		}},
		{name: "HandleGet_Error", method: http.MethodGet, route: "/api/groups/%d", groupID: 1, err: errors.New("some error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetGroupByID", int32(1)).Return(&groupDomain.Group{}, errors.New("some error"))
		}},
		{name: "HandleGet_InvalidID", method: http.MethodGet, route: "/api/groups/invalid", groupID: 0, err: nil, expStatus: http.StatusBadRequest, mockSetupFunc: func(ms *MockService) {}},
		{name: "HandleGetAll_Error", method: http.MethodGet, route: "/api/groups", groupID: 0, err: errors.New("some error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetAllGroups").Return([]*groupDomain.Group{}, errors.New("some error"))
		}},
		{name: "NonExistentRoute", method: http.MethodGet, route: "/api/non_existent_route", groupID: 0, err: nil, expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {}},
		{name: "NotAllowedMethod", method: http.MethodPost, route: "/api/groups/%d", groupID: 1, err: &utils.MethodNotAllowedError{Method: "POST", Route: "/api/groups/%d"}, expStatus: http.StatusMethodNotAllowed, mockSetupFunc: func(ms *MockService) {}},
		{name: "ServiceError", method: http.MethodGet, route: "/api/groups/%d", groupID: 1, err: errors.New("internal server error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetGroupByID", int32(1)).Return(nil, errors.New("some other error"))
		}},
		{name: "NoResults", method: http.MethodGet, route: "/api/groups/%d", groupID: 999, err: &utils.NoResultsForParameterError{ParameterName: "group_id", ID: "999"}, expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {
			ms.On("GetGroupByID", int32(999)).Return(nil, pgx.ErrNoRows)
		}},
		{name: "UnexpectedError", method: http.MethodGet, route: "/api/groups/%d", groupID: 1, err: errors.New("unexpected error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetGroupByID", int32(1)).Return(nil, errors.New("unexpected error"))
		}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			mockService := &MockService{}
			service := groupDomain.NewService(mockService)
			handler := NewGroupHandler(service)
			tt.mockSetupFunc(mockService)

			route := tt.route
			if tt.groupID != 0 {
				route = fmt.Sprintf(tt.route, tt.groupID)
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
