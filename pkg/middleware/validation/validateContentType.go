package validation

import (
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"net/http"
)

func ValidateContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request media type is supported
		if r.Header.Get("Content-Type") != "application/json" {
			httpError := &utils.InvalidParameterError{ParameterName: "Content-Type"}
			utils.HandleHTTPError(w, httpError)
			return
		}

		// If the media type is supported, call the next handler
		next.ServeHTTP(w, r)
	})
}
