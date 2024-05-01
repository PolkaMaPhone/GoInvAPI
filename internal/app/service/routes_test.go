package service

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAPIHandler struct {
	mock.Mock
}

func (m *MockAPIHandler) HandleGetItem(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func TestGetItemRoute(t *testing.T) {
	mockHandler := new(MockAPIHandler)
	mockHandler.On("HandleGetItem", mock.Anything, mock.Anything).Return()

	r := mux.NewRouter()
	r.HandleFunc("/items/{item_id}", mockHandler.HandleGetItem).Methods("GET")

	req, _ := http.NewRequest("GET", "/items/1", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	mockHandler.AssertExpectations(t)
}
