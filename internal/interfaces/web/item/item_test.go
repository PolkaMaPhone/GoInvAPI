package item

import (
	"encoding/json"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockItemHandler struct {
	mock.Mock
}

func (m *MockItemHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockItemHandler) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/items/{item_id}", m.HandleGet).Methods("GET")
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

func TestHandleGetItem(t *testing.T) {
	config, err := dbconn.LoadConfigFile()
	if err != nil {
		t.Fatalf("Unable to load configuration: %v\n", err)
	}
	db := &dbconn.PgxDB{}
	_, err = dbconn.New(config, db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	handler := NewItemHandler(db.Pool)

	req, err := http.NewRequest("GET", "/items/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items/{item_id}", handler.HandleGet)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Unmarshal the response body into a map
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Could not unmarshal response body: %v", err)
	}

	// Check for the existence of the timestamp fields
	if _, ok := response["CreatedAt"]; !ok {
		t.Errorf("Expected 'CreatedAt' field in response body")
	}
	if _, ok := response["UpdatedAt"]; !ok {
		t.Errorf("Expected 'UpdatedAt' field in response body")
	}

	// Check the other fields as before
	expectedItem := map[string]interface{}{
		"ItemID":      1.0, // json.Unmarshal converts integers to floats
		"Name":        "Sample Item 111",
		"Description": "Description for Sample Item 1-1-1",
		"CategoryID":  1.0,
		"GroupID":     1.0,
		"LocationID":  1.0,
		"IsStored":    false,
	}
	for key, expectedValue := range expectedItem {
		if response[key] != expectedValue {
			t.Errorf("handler returned unexpected %s: got %v want %v",
				key, response[key], expectedValue)
		}
	}
}
