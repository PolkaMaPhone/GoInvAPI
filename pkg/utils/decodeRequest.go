package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeFromRequest(w http.ResponseWriter, r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		HandleHTTPError(w, err)
		return err
	}
	return nil
}
