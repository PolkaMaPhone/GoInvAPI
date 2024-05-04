package customRouter

import (
	"context"
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type CustomRouter struct {
	*mux.Router
	limiter *rate.Limiter
}

func NewRouter() *CustomRouter {
	r := mux.NewRouter()
	limiter := rate.NewLimiter(1, 5) // Adjust to your needs
	customRouter := &CustomRouter{r, limiter}
	customRouter.NotFoundHandler = http.HandlerFunc(customRouter.handleNotFound)
	customRouter.MethodNotAllowedHandler = http.HandlerFunc(customRouter.handleMethodNotAllowed)
	return customRouter
}

func (cr *CustomRouter) handleNotFound(w http.ResponseWriter, r *http.Request) {
	middleware.ErrorLogger.Printf("Invalid route accessed: %s", r.URL.Path)
	http.Error(w, "Invalid route. Please check the URL and try again.", http.StatusNotFound)
}

func (cr *CustomRouter) handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	middleware.ErrorLogger.Printf("Method not allowed: %s", r.Method)
	http.Error(w, "Method not allowed. Please check the HTTP method and try again.", http.StatusMethodNotAllowed)
}

func (cr *CustomRouter) handleRequestTimeout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			middleware.ErrorLogger.Printf("Request timeout: %s", r.URL.Path)
			http.Error(w, "Request timeout. Please try again later.", http.StatusRequestTimeout)
		}
	case <-time.After(5 * time.Second): // Replace with your desired timeout duration
		// Continue processing the request
	}
}

func (cr *CustomRouter) handleUnsupportedMediaType(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" { // Replace with your supported media types
		middleware.ErrorLogger.Printf("Unsupported media type: %s", r.Header.Get("Content-Type"))
		http.Error(w, "Unsupported media type. Please check the 'Content-Type' header and try again.", http.StatusUnsupportedMediaType)
	}
}

func (cr *CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cr.limiter.Allow() == false {
		http.Error(w, "Too Many Requests - try again later", http.StatusTooManyRequests)
		return
	}

	cr.Router.ServeHTTP(w, r)
}
