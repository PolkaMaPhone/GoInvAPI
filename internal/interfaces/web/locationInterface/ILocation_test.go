package locationInterface

import (
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/locationDomain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockLocationHandler struct {
	mock.Mock
}

type MockService struct {
	mock.Mock
}

func (m *MockLocationHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockLocationHandler) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/locations/{location_id}", m.HandleGet).Methods("GET")
}

func (m *MockService) GetLocationByID(id int32) (*locationDomain.Location, error) {
	args := m.Called(id)
	return args.Get(0).(*locationDomain.Location), args.Error(1)
}

func (m *MockService) GetAllLocations() ([]*locationDomain.Location, error) {
	args := m.Called()
	return args.Get(0).([]*locationDomain.Location), args.Error(1)
}

func TestGetLocationRoute(t *testing.T) {
	mockHandler := new(MockLocationHandler)
	mockHandler.On("HandleGet", mock.Anything, mock.AnythingOfType("*http.Request")).Return()

	r := mux.NewRouter()
	mockHandler.HandleRoutes(r)

	req, _ := http.NewRequest("GET", "/locations/1", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	mockHandler.AssertCalled(t, "HandleGet", rr, mock.AnythingOfType("*http.Request"))
}

func TestHandleGet_Error(t *testing.T) {
	mockService := new(MockService)
	mockLocationService := locationDomain.NewService(mockService)
	handler := NewLocationHandler(mockLocationService)

	mockService.On("GetLocationByID", int32(1)).Return(&locationDomain.Location{}, errors.New("some error"))

	req, _ := http.NewRequest("GET", "/locations/1", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/locations/{location_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestHandleGet_InvalidID(t *testing.T) {
	mockService := new(MockService)
	mockLocationService := locationDomain.NewService(mockService)
	handler := NewLocationHandler(mockLocationService)

	req, _ := http.NewRequest("GET", "/locations/invalid", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/locations/{location_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandleGetAll_Error(t *testing.T) {
	mockService := new(MockService)
	mockLocationService := locationDomain.NewService(mockService)
	handler := NewLocationHandler(mockLocationService)
	mockService.On("GetAllLocations").Return([]*locationDomain.Location{}, errors.New("some error"))

	req, _ := http.NewRequest("GET", "/locations", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/locations", handler.HandleGetAll)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}
