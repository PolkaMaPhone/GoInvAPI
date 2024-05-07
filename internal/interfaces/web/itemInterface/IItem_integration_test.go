package itemInterface

import (
	"bytes"
	"net/http"
	"testing"
)

func TestItemIntegration(t *testing.T) {
	testCases := []struct {
		name   string
		method string
		route  string
		status int
	}{
		{name: "HandleGet", method: http.MethodGet, route: "/api/items/1", status: http.StatusOK},
		{name: "HandleGetAll", method: http.MethodGet, route: "/api/items", status: http.StatusOK},
		{name: "HandleGet_NonExistentItem", method: http.MethodGet, route: "/api/items/9999", status: http.StatusNotFound},
		{name: "HandleGet_InvalidID", method: http.MethodGet, route: "/api/items/invalid", status: http.StatusBadRequest},
		{name: "NotAllowedMethod", method: http.MethodPost, route: "/api/items/1", status: http.StatusMethodNotAllowed},
	}
	server := initializeItemTestServer()
	defer server.Close()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.NewRequest(tt.method, server.URL+tt.route, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			client := &http.Client{}
			response, err := client.Do(resp)
			if err != nil {
				t.Fatalf("could not send request: %v", err)
			}

			if response.StatusCode != tt.status {
				t.Errorf("expected status %d, got %d", tt.status, response.StatusCode)
			}
		})
	}
}

func TestItemCRUD(t *testing.T) {
	testCases := []struct {
		name                         string
		itemJSON                     []byte
		updatedItemJSON              []byte
		expectedCreateStatus         int
		expectedGetStatus            int
		expectedUpdateStatus         int
		expectedDeleteStatus         int
		expectedGetAfterDeleteStatus int
	}{
		{
			name:                         "Test case 1",
			itemJSON:                     []byte(`{"name":"item1","description":"description1"}`),
			updatedItemJSON:              []byte(`{"name":"updatedItem1","description":"updatedDescription1"}`),
			expectedCreateStatus:         http.StatusCreated,
			expectedGetStatus:            http.StatusOK,
			expectedUpdateStatus:         http.StatusOK,
			expectedDeleteStatus:         http.StatusOK,
			expectedGetAfterDeleteStatus: http.StatusNotFound,
		},
	}
	server := initializeItemTestServer()
	defer server.Close()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new item
			resp, err := http.Post(server.URL+"/api/items", "application/json", bytes.NewBuffer(tt.itemJSON))
			if err != nil {
				t.Fatalf("could not create item: %v", err)
			}
			if resp.StatusCode != tt.expectedCreateStatus {
				t.Errorf("expected status %d, got %d", tt.expectedCreateStatus, resp.StatusCode)
			}

			// Get the created item
			resp, err = http.Get(server.URL + "/api/items/1")
			if err != nil {
				t.Fatalf("could not get item: %v", err)
			}
			if resp.StatusCode != tt.expectedGetStatus {
				t.Errorf("expected status %d, got %d", tt.expectedGetStatus, resp.StatusCode)
			}

			// Update the item
			client := &http.Client{}
			req, err := http.NewRequest(http.MethodPut, server.URL+"/api/items/1", bytes.NewBuffer(tt.updatedItemJSON))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			resp, err = client.Do(req)
			if err != nil {
				t.Fatalf("could not update item: %v", err)
			}
			if resp.StatusCode != tt.expectedUpdateStatus {
				t.Errorf("expected status %d, got %d", tt.expectedUpdateStatus, resp.StatusCode)
			}

			// Delete the item
			req, err = http.NewRequest(http.MethodDelete, server.URL+"/api/items/1", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			resp, err = client.Do(req)
			if err != nil {
				t.Fatalf("could not delete item: %v", err)
			}
			if resp.StatusCode != tt.expectedDeleteStatus {
				t.Errorf("expected status %d, got %d", tt.expectedDeleteStatus, resp.StatusCode)
			}

			// Try to get the deleted item
			resp, err = http.Get(server.URL + "/api/items/1")
			if err != nil {
				t.Fatalf("could not get item: %v", err)
			}
			if resp.StatusCode != tt.expectedGetAfterDeleteStatus {
				t.Errorf("expected status %d, got %d", tt.expectedGetAfterDeleteStatus, resp.StatusCode)
			}
		})
	}
}
