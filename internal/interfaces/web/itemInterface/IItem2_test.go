package itemInterface

import (
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestCase struct {
	name           string
	route          string
	expectedErr    error
	expectedId     int32
	expectedStatus int
	expectedBody   string
}

type MockSetupFunc func(ms *MockService, tc TestCase)

func RunHandlerTests(t *testing.T, methodName string, testCases []TestCase, mockSetupFunc MockSetupFunc) {
	// Map of method names to handler methods
	handlerMethods := map[string]func(*Handler, *mux.Router){
		"GetItemByID": func(h *Handler, router *mux.Router) {
			router.HandleFunc("/items/{item_id}", h.HandleGet).Methods("GET")
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.route, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			ms := new(MockService)
			mockSetupFunc(ms, tt)
			s := itemDomain.NewService(ms)
			h := NewItemHandler(s)

			// Call the appropriate handler method based on the method name
			if method, ok := handlerMethods[methodName]; ok {
				method(h, router)
			} else {
				t.Fatalf("Invalid method name: %s", methodName)
			}

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(tt.expectedBody) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestHandleGet(t *testing.T) {
	testCases := []TestCase{
		// Successful test cases
		{name: "getItem_ValidID1", route: "/items/1", expectedId: 1, expectedStatus: http.StatusOK, expectedBody: `{"ItemID":1,"Name":"","Description":null,"CategoryID":null,"GroupID":null,"LocationID":null,"IsStored":null,"CreatedAt":null,"UpdatedAt":null}`},
		{name: "getItem_ValidID2", route: "/items/2", expectedId: 2, expectedStatus: http.StatusOK, expectedBody: `{"ItemID":2,"Name":"","Description":null,"CategoryID":null,"GroupID":null,"LocationID":null,"IsStored":null,"CreatedAt":null,"UpdatedAt":null}`},

		// Failed test cases
		{name: "getItem_InvalidID_NotFound", route: "/items/999", expectedId: 999, expectedStatus: http.StatusNotFound, expectedBody: fmt.Sprintf(utils.HTTPErrorMessages["NoResultsForParameter"], "item_id", "999"), expectedErr: &utils.NoResultsForParameterError{ParameterName: "item_id", ID: "999", StatusCode: http.StatusNotFound}},
		{name: "getItem_InvalidID_Format", route: "/items/invalid", expectedId: 0, expectedStatus: http.StatusBadRequest, expectedBody: fmt.Sprintf(utils.HTTPErrorMessages["InvalidParameter"], "item_id"), expectedErr: &utils.InvalidParameterError{ParameterName: "item_id"}},
		{name: "getItem_ValidID_DatabaseError", route: "/items/3", expectedId: 3, expectedStatus: http.StatusInternalServerError, expectedBody: utils.HTTPErrorMessages[utils.ServerError], expectedErr: &utils.ServerErrorType{}},
	}
	mockSetupFunc := func(ms *MockService, tc TestCase) {
		// If the expected error is not nil and is a database error, return nil for the item
		if _, ok := tc.expectedErr.(*utils.ServerErrorType); ok {
			ms.On("GetItemByID", tc.expectedId).Return(nil, tc.expectedErr)
		} else if _, ok := tc.expectedErr.(*utils.NoResultsForParameterError); ok {
			// Ensure a valid item is returned for the test case
			ms.On("GetItemByID", tc.expectedId).Return(nil, pgx.ErrNoRows)
		} else {
			ms.On("GetItemByID", tc.expectedId).Return(&itemDomain.Item{ItemID: tc.expectedId}, tc.expectedErr)
		}
	}

	RunHandlerTests(t, "GetItemByID", testCases, mockSetupFunc)
}
