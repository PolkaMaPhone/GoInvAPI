package validation

import (
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"net/http"
)

func ValidateMethod(allowedMethods ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			methodAllowed := false
			for _, method := range allowedMethods {
				if r.Method == method {
					methodAllowed = true
					break
				}
			}

			if !methodAllowed {
				httpError := &utils.MethodNotAllowedError{Method: r.Method, Route: r.URL.Path}
				utils.HandleHTTPError(w, httpError)
				return
			}

			// If the method is allowed, call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
