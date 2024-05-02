package statusInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct{}

func NewStatusHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.LoggingMiddleware("INFO"))
	apiRouter.HandleFunc("/status", h.HandleStatus).Methods("GET")
}

func (h *Handler) HandleStatus(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Server is up and running"))
	if err != nil {
		return
	}
}
