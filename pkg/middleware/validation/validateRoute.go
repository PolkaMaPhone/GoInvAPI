package validation

import (
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ValidateRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the accessed route
		logging.InfoLogger.Printf("Accessed route: %s", r.URL.Path)

		// If the route does not exist, respond with an error
		if r.Context().Value(chi.RouteCtxKey) == nil {
			httpError := &utils.InvalidRouteError{Route: r.URL.Path}
			utils.HandleHTTPError(w, httpError)
			return
		}

		// If the route exists, call the next handler
		next.ServeHTTP(w, r)
	})
}
