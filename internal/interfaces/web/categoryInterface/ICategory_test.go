package categoryInterface

import (
	"errors"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/categoryDomain"
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

func (m *MockService) GetCategoryByID(id int32) (*categoryDomain.Category, error) {
	args := m.Called(id)
	category, ok := args.Get(0).(*categoryDomain.Category)
	if !ok {
		return nil, args.Error(1)
	}
	return category, args.Error(1)
}

func (m *MockService) GetAllCategories() ([]*categoryDomain.Category, error) {
	args := m.Called()
	return args.Get(0).([]*categoryDomain.Category), args.Error(1)
}

func TestHandleRoutes(t *testing.T) {
	testCases := []struct {
		name          string
		method        string
		route         string
		categoryID    int
		err           error
		expStatus     int
		mockSetupFunc func(ms *MockService)
	}{
		{name: "HandleGet", method: http.MethodGet, route: "/categories/%d", categoryID: 1, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetCategoryByID", int32(1)).Return(&categoryDomain.Category{CategoryID: 1}, nil)
		}},
		{name: "HandleGet_Error", method: http.MethodGet, route: "/categories/%d", categoryID: 1, err: errors.New("some error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetCategoryByID", int32(1)).Return(&categoryDomain.Category{}, errors.New("some error"))
		}},
		{name: "HandleGet_InvalidID", method: http.MethodGet, route: "/categories/invalid", categoryID: 0, err: nil, expStatus: http.StatusBadRequest, mockSetupFunc: func(ms *MockService) {}},
		{name: "HandleGetAll_Error", method: http.MethodGet, route: "/categories", categoryID: 0, err: errors.New("some error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetAllCategories").Return([]*categoryDomain.Category{}, errors.New("some error"))
		}},
		{name: "NonExistentRoute", method: http.MethodGet, route: "/non_existent_route", categoryID: 0, err: nil, expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {}},
		{name: "NotAllowedMethod", method: http.MethodPost, route: "/categories/%d", categoryID: 1, err: &utils.MethodNotAllowedError{Method: "POST", Route: "/api/categories/%d"}, expStatus: http.StatusMethodNotAllowed, mockSetupFunc: func(ms *MockService) {}},
		{name: "ServiceError", method: http.MethodGet, route: "/categories/%d", categoryID: 1, err: errors.New("internal server error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetCategoryByID", int32(1)).Return(nil, errors.New("some other error"))
		}},
		{name: "NoResults", method: http.MethodGet, route: "/categories/%d", categoryID: 999, err: &utils.NoResultsForParameterError{ParameterName: "category_id", ID: "999", StatusCode: http.StatusNotFound}, expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {
			ms.On("GetCategoryByID", int32(999)).Return(nil, pgx.ErrNoRows)
		}},
		{name: "UnexpectedError", method: http.MethodGet, route: "/categories/%d", categoryID: 1, err: errors.New("unexpected error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetCategoryByID", int32(1)).Return(nil, errors.New("unexpected error"))
		}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			mockService := &MockService{}
			service := categoryDomain.NewService(mockService)
			handler := NewCategoryHandler(service)
			tt.mockSetupFunc(mockService)

			route := tt.route
			if tt.categoryID != 0 {
				route = fmt.Sprintf(tt.route, tt.categoryID)
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
