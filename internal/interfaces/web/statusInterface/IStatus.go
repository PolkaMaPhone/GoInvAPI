package statusInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/errorMiddleware"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validationMiddleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct{}

func NewStatusHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleRoutes(apiRouter *customRouter.CustomRouter) {
	r := chi.NewRouter()
	r.Use(validationMiddleware.ValidateMethod(http.MethodGet))
	r.Get("/", errorMiddleware.WithErrorHandling(h.HandleStatus))
	apiRouter.Mount(apiRouter.GetFullPath("/status"), r)
}

func (h *Handler) HandleStatus(w http.ResponseWriter, _ *http.Request) error {
	_, err := w.Write([]byte("Server is up and running"))
	if err != nil {
		return err
	}
	return nil
}
