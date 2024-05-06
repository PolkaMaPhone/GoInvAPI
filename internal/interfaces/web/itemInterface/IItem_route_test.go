package itemInterface

import (
	"errors"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
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
		itemID        int
		err           error
		expStatus     int
		mockSetupFunc func(ms *MockService)
	}{
		{name: "HandleGet", method: http.MethodGet, route: "/items/%d", itemID: 1, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetItemByID", int32(1)).Return(&itemDomain.Item{ItemID: 1}, nil)
		}},
		{name: "HandleGetWithCategory", method: http.MethodGet, route: "/items/%d/with_category", itemID: 1, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetItemByIDWithCategory", int32(1)).Return(&dto.ItemWithCategory{ItemID: 1}, nil)
		}},
		{name: "HandleGetWithGroup", method: http.MethodGet, route: "/items/%d/with_group", itemID: 1, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetItemByIDWithGroup", int32(1)).Return(&dto.ItemWithGroup{ItemID: 1}, nil)
		}},
		{name: "HandleGetWithGroupAndCategory", method: http.MethodGet, route: "/items/%d/with_group_and_category", itemID: 1, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetItemByIDWithGroupAndCategory", int32(1)).Return(&dto.ItemWithGroupAndCategory{ItemID: 1}, nil)
		}},
		{name: "HandleGetAll", method: http.MethodGet, route: "/items", itemID: 0, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetAllItems").Return([]*itemDomain.Item{{ItemID: int32(1)}, {ItemID: int32(2)}}, nil)
		}},
		{name: "HandleGetAllWithCategories", method: http.MethodGet, route: "/items_with_category", itemID: 0, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetAllItemsWithCategories").Return([]*dto.ItemWithCategory{{ItemID: int32(1)}, {ItemID: int32(2)}}, nil)
		}},
		{name: "HandleGetAllWithGroups", method: http.MethodGet, route: "/items_with_group", itemID: 0, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetAllItemsWithGroups").Return([]*dto.ItemWithGroup{{ItemID: int32(1)}, {ItemID: int32(2)}}, nil)
		}},
		{name: "HandleGetAllWithGroupsAndCategories", method: http.MethodGet, route: "/items_with_group_and_category", itemID: 0, err: nil, expStatus: http.StatusOK, mockSetupFunc: func(ms *MockService) {
			ms.On("GetAllItemsWithGroupsAndCategories").Return([]*dto.ItemWithGroupAndCategory{{ItemID: int32(1)}, {ItemID: int32(2)}}, nil)
		}},
		{name: "NonExistentRoute", method: http.MethodGet, route: "/non_existent_route", itemID: 0, err: nil, expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {}},
		{name: "NotAllowedMethod", method: http.MethodPost, route: "/items/%d", itemID: 1, err: &utils.MethodNotAllowedError{Method: "POST", Route: "/api/items/%d"}, expStatus: http.StatusMethodNotAllowed, mockSetupFunc: func(ms *MockService) {}},
		{name: "ServiceError", method: http.MethodGet, route: "/items/%d", itemID: 1, err: errors.New("internal server error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetItemByID", int32(1)).Return(nil, errors.New("some other error"))
		}},
		{name: "NoResults", method: http.MethodGet, route: "/items/%d", itemID: 999, err: &utils.NoResultsForParameterError{ParameterName: "item_id", ID: "999", StatusCode: http.StatusNotFound}, expStatus: http.StatusNotFound, mockSetupFunc: func(ms *MockService) {
			ms.On("GetItemByID", int32(999)).Return(nil, pgx.ErrNoRows)
		}},
		{name: "UnexpectedError", method: http.MethodGet, route: "/items/%d", itemID: 1, err: errors.New("unexpected error"), expStatus: http.StatusInternalServerError, mockSetupFunc: func(ms *MockService) {
			ms.On("GetItemByID", int32(1)).Return(nil, errors.New("unexpected error"))
		}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			mockService := &MockService{}
			service := itemDomain.NewService(mockService)
			handler := NewItemHandler(service)
			tt.mockSetupFunc(mockService)

			route := tt.route
			if tt.itemID != 0 {
				route = fmt.Sprintf(tt.route, tt.itemID)
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
