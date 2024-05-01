package apihandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetItem(t *testing.T) {
	handler := NewAPIHandler()

	req, err := http.NewRequest("GET", "/items/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items/{item_id}", handler.HandleGetItem)

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

func TestHandleStatus(t *testing.T) {
	handler := NewAPIHandler()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.HandleStatus(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"port":"","status":"server running"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
