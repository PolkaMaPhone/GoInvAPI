package itemInterface

import (
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
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

func (m *MockService) GetItemByIDWithGroup(id int32) (*dto.ItemWithGroup, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ItemWithGroup), args.Error(1)
}

func (m *MockService) GetItemByIDWithGroupAndCategory(id int32) (*dto.ItemWithGroupAndCategory, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ItemWithGroupAndCategory), args.Error(1)
}

func (m *MockService) GetAllItemsWithGroups() ([]*dto.ItemWithGroup, error) {
	args := m.Called()
	return args.Get(0).([]*dto.ItemWithGroup), args.Error(1)
}

func (m *MockService) GetAllItemsWithGroupsAndCategories() ([]*dto.ItemWithGroupAndCategory, error) {
	args := m.Called()
	return args.Get(0).([]*dto.ItemWithGroupAndCategory), args.Error(1)
}

func (m *MockItemHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockItemHandler) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/items/{item_id}", m.HandleGet).Methods("GET")
}

func (m *MockService) GetItemByID(id int32) (*itemDomain.Item, error) {
	args := m.Called(id)
	item, ok := args.Get(0).(*itemDomain.Item)
	if !ok {
		return nil, args.Error(1)
	}
	return item, args.Error(1)
}

func (m *MockService) GetAllItems() ([]*itemDomain.Item, error) {
	args := m.Called()
	return args.Get(0).([]*itemDomain.Item), args.Error(1)
}

func (m *MockService) GetItemByIDWithCategory(id int32) (*dto.ItemWithCategory, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ItemWithCategory), args.Error(1)
}

func (m *MockService) GetAllItemsWithCategories() ([]*dto.ItemWithCategory, error) {
	args := m.Called()
	return args.Get(0).([]*dto.ItemWithCategory), args.Error(1)
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

func TestGetCategoryRoute(t *testing.T) {
	mockHandler := new(MockItemHandler)
	mockHandler.On("HandleGet", mock.Anything, mock.AnythingOfType("*http.Request")).Return()

	r := mux.NewRouter()
	mockHandler.HandleRoutes(r)

	req, _ := http.NewRequest("GET", "/items/1/category", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	mockHandler.AssertNotCalled(t, "HandleGet", rr, mock.AnythingOfType("*http.Request"))

}

func TestHandler_HandleGetAllWithCategory(t *testing.T) {
	mockService := new(MockService)
	mockItemService := itemDomain.NewService(mockService)
	handler := NewItemHandler(mockItemService)
	mockService.On("GetAllItemsWithCategories").Return([]*dto.ItemWithCategory{}, errors.New("some error"))

	req, _ := http.NewRequest("GET", "/items_with_category", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items_with_category", handler.HandleGetAllWithCategories)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}
