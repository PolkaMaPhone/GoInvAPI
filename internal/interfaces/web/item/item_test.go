package item

import (
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/item"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockItemHandler struct {
	mock.Mock
}

type MockService struct {
	mock.Mock
}

func (m *MockItemHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockItemHandler) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/items/{item_id}", m.HandleGet).Methods("GET")
}

func (m *MockService) GetItemByID(id int32) (*item.Item, error) {
	args := m.Called(id)
	return args.Get(0).(*item.Item), args.Error(1)
}

func (m *MockService) GetAllItems() ([]*item.Item, error) {
	args := m.Called()
	return args.Get(0).([]*item.Item), args.Error(1)
}

func TestGetItemRoute(t *testing.T) {
	mockHandler := new(MockItemHandler)
	mockHandler.On("HandleGet", mock.Anything, mock.AnythingOfType("*http.Request")).Return()

	r := mux.NewRouter()
	mockHandler.HandleRoutes(r)

	req, _ := http.NewRequest("GET", "/items/1", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	mockHandler.AssertCalled(t, "HandleGet", rr, mock.AnythingOfType("*http.Request"))
}

func TestHandleGet_Error(t *testing.T) {
	mockService := new(MockService)
	mockItemService := item.NewService(mockService)
	handler := NewItemHandler(mockItemService)

	mockService.On("GetItemByID", int32(1)).Return(&item.Item{}, errors.New("some error"))

	req, _ := http.NewRequest("GET", "/items/1", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items/{item_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestHandleGet_InvalidID(t *testing.T) {
	mockService := new(MockService)
	mockItemService := item.NewService(mockService)
	handler := NewItemHandler(mockItemService)

	req, _ := http.NewRequest("GET", "/items/invalid", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items/{item_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandleGetAll_Error(t *testing.T) {
	mockService := new(MockService)
	mockItemService := item.NewService(mockService)
	handler := NewItemHandler(mockItemService)
	mockService.On("GetAllItems").Return([]*item.Item{}, errors.New("some error"))

	req, _ := http.NewRequest("GET", "/items", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items", handler.HandleGetAll)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}
