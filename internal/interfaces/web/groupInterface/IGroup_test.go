package groupInterface

import (
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/groupDomain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockGroupHandler struct {
	mock.Mock
}

func (m *MockGroupHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockGroupHandler) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/groups/{group_id}", m.HandleGet).Methods("GET")
}

type MockService struct {
	mock.Mock
}

func (m *MockService) GetGroupByID(id int32) (*groupDomain.Group, error) {
	args := m.Called(id)
	return args.Get(0).(*groupDomain.Group), args.Error(1)
}

func (m *MockService) GetAllGroups() ([]*groupDomain.Group, error) {
	args := m.Called()
	return args.Get(0).([]*groupDomain.Group), args.Error(1)
}

func generateRequest(t *testing.T, method, url string) (*httptest.ResponseRecorder, *http.Request) {
	req, err := http.NewRequest(method, url, nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	return rr, req
}

func TestGetGroupRoute(t *testing.T) {
	mockHandler := new(MockGroupHandler)
	mockHandler.On("HandleGet", mock.Anything, mock.AnythingOfType("*http.Request")).Return()

	r := mux.NewRouter()
	mockHandler.HandleRoutes(r)

	rr, req := generateRequest(t, "GET", "/groups/1")
	r.ServeHTTP(rr, req)

	mockHandler.AssertCalled(t, "HandleGet", rr, mock.AnythingOfType("*http.Request"))
}

func TestHandleGet_Error(t *testing.T) {
	mockService := new(MockService)
	mockGroupService := groupDomain.NewService(mockService)
	handler := NewGroupHandler(mockGroupService)

	mockService.On("GetGroupByID", int32(1)).Return(&groupDomain.Group{}, errors.New("some error"))

	rr, req := generateRequest(t, "GET", "/groups/1")
	router := mux.NewRouter()
	router.HandleFunc("/groups/{group_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestHandleGet_InvalidID(t *testing.T) {
	mockService := new(MockService)
	mockGroupService := groupDomain.NewService(mockService)
	handler := NewGroupHandler(mockGroupService)

	rr, req := generateRequest(t, "GET", "/groups/invalid")
	router := mux.NewRouter()
	router.HandleFunc("/groups/{group_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandleGetAll_Error(t *testing.T) {
	mockService := new(MockService)
	mockGroupService := groupDomain.NewService(mockService)
	handler := NewGroupHandler(mockGroupService)
	mockService.On("GetAllGroups").Return([]*groupDomain.Group{}, errors.New("some error"))

	rr, req := generateRequest(t, "GET", "/groups")
	router := mux.NewRouter()
	router.HandleFunc("/groups", handler.HandleGetAll)
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}
