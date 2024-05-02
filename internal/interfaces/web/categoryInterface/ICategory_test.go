package categoryInterface

import (
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/categoryDomain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockCategoryHandler struct {
	mock.Mock
}

type MockService struct {
	mock.Mock
}

func (m *MockCategoryHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockCategoryHandler) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/categories/{category_id}", m.HandleGet).Methods("GET")
}

func (m *MockService) GetCategoryByID(id int32) (*categoryDomain.Category, error) {
	args := m.Called(id)
	return args.Get(0).(*categoryDomain.Category), args.Error(1)
}

func (m *MockService) GetAllCategories() ([]*categoryDomain.Category, error) {
	args := m.Called()
	return args.Get(0).([]*categoryDomain.Category), args.Error(1)
}

func TestGetCategoryRoute(t *testing.T) {
	mockHandler := new(MockCategoryHandler)
	mockHandler.On("HandleGet", mock.Anything, mock.AnythingOfType("*http.Request")).Return()

	r := mux.NewRouter()
	mockHandler.HandleRoutes(r)

	req, _ := http.NewRequest("GET", "/categories/1", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	mockHandler.AssertCalled(t, "HandleGet", rr, mock.AnythingOfType("*http.Request"))
}

func TestHandleGet_Error(t *testing.T) {
	mockService := new(MockService)
	mockCategoryService := categoryDomain.NewService(mockService)
	handler := NewCategoryHandler(mockCategoryService)

	mockService.On("GetCategoryByID", int32(1)).Return(&categoryDomain.Category{}, errors.New("some error"))

	req, _ := http.NewRequest("GET", "/categories/1", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/categories/{category_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestHandleGet_InvalidID(t *testing.T) {
	mockService := new(MockService)
	mockCategoryService := categoryDomain.NewService(mockService)
	handler := NewCategoryHandler(mockCategoryService)

	req, _ := http.NewRequest("GET", "/categories/invalid", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/categories/{category_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandleGetAll_Error(t *testing.T) {
	mockService := new(MockService)
	mockCategoryService := categoryDomain.NewService(mockService)
	handler := NewCategoryHandler(mockCategoryService)
	mockService.On("GetAllCategories").Return([]*categoryDomain.Category{}, errors.New("some error"))

	req, _ := http.NewRequest("GET", "/categories", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/categories", handler.HandleGetAll)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}
