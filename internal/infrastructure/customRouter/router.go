package customRouter

import (
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/errorMiddleware"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validationMiddleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/time/rate"
	"net/http"
)

type CustomRouter struct {
	*chi.Mux
	limiter *rate.Limiter
	prefix  string
}

func NewDefaultRouter() *CustomRouter {
	return NewRouter("INFO", "/api")
}

func NewRouter(logLevel string, prefix string) *CustomRouter {
	r := chi.NewRouter()
	limiter := rate.NewLimiter(1, 5)

	// Use middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(logging.LogRequestDuration(logLevel))
	r.Use(validationMiddleware.ValidateRoute)
	r.Use(validationMiddleware.ValidateContentType)
	r.Use(WithErrorHandlingMiddleware)

	// Create a sub-router
	subRouter := chi.NewRouter()
	r.Mount(prefix, subRouter)

	customRouter := &CustomRouter{subRouter, limiter, prefix}
	return customRouter
}

func WithErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMiddleware.WithErrorHandling(func(w http.ResponseWriter, r *http.Request) error {
			next.ServeHTTP(w, r)
			return nil
		})(w, r)
	})
}

func (cr *CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cr.limiter.Allow() == false {
		http.Error(w, "Too Many Requests - try again later", http.StatusTooManyRequests)
		return
	}

	cr.Mux.ServeHTTP(w, r)
}

func (cr *CustomRouter) GetFullPath(path string) string {
	return cr.prefix + path
}
