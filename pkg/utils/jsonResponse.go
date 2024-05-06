// Package utils Description: This file contains the function to write json response format.
package utils

import (
	"encoding/json"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"net/http"
)

// RespondWithJSON writes json response format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		logging.ErrorLogger.Printf("Error writing response: %v", err)
	}
}
