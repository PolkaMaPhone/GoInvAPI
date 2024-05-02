package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testPayload struct {
	Message string `json:"message"`
}

func TestRespondWithJSON(t *testing.T) {
	tests := []struct {
		name           string
		code           int
		payload        interface{}
		expectedBody   string
		expectedStatus int
	}{
		{
			name:           "Successful response",
			code:           http.StatusOK,
			payload:        testPayload{"Success"},
			expectedBody:   "{\"message\":\"Success\"}",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Error response",
			code:           http.StatusInternalServerError,
			payload:        testPayload{"Error"},
			expectedBody:   "{\"message\":\"Error\"}",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Empty response",
			code:           http.StatusOK,
			payload:        testPayload{},
			expectedBody:   "{\"message\":\"\"}",
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			recorder := httptest.NewRecorder()

			RespondWithJSON(recorder, test.code, test.payload)

			result := recorder.Result()

			_, err = json.Marshal(test.payload)
			if err != nil {
				t.Fatalf("Could not marshal payload: %v", err)
			}

			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(result.Body)
			if err != nil {
				return
			}
			newStr := buf.String()

			if newStr != test.expectedBody {
				t.Errorf("Expected body %v but got %v", []byte(test.expectedBody), []byte(newStr))
			}

			if result.StatusCode != test.expectedStatus {
				t.Errorf("Expected status %v but got %v", test.expectedStatus, result.StatusCode)
			}
		})
	}
}
