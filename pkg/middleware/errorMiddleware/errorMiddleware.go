package errorMiddleware

import (
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"log"
	"net/http"
)

type HTTPError interface {
	HTTPStatus() int
	Error() string
}

func WithErrorHandling(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			// Log the error
			log.Printf("Error handling request: %v", err)

			// TODO - Maybe go crazy and implement some metrics here
			// Increment error metric (pseudo-code)
			// metrics.Increment("http_errors")

			// type assertion is fine here
			if httpErr, ok := err.(HTTPError); ok {
				// Send a custom error response
				http.Error(w, httpErr.Error(), httpErr.HTTPStatus())
			} else {
				// Send a generic error response
				utils.HandleHTTPError(w, err)
			}
		}
	}
}
