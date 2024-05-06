package utils

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func GetIDFromRequest(w http.ResponseWriter, r *http.Request, parameterName string) (int32, error) {
	idStr := chi.URLParam(r, parameterName)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleHTTPError(w, &InvalidParameterError{ParameterName: parameterName})
		return 0, err
	}
	return int32(id), nil
}
