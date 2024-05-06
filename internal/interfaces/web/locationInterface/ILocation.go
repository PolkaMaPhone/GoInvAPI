package locationInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/locationDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validation"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const idParameterName = "location_id"

type Handler struct {
	service *locationDomain.Service
}

func NewLocationHandler(s *locationDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(apiRouter *customRouter.CustomRouter) {
	r := chi.NewRouter()
	r.Use(validation.ValidateMethod(http.MethodGet))
	r.Get("/", h.HandleGet)
	apiRouter.Mount(apiRouter.GetFullPath("/locations/{location_id}"), r)

	r = chi.NewRouter()
	r.Use(validation.ValidateMethod(http.MethodGet))
	r.Get("/", h.HandleGetAll)
	apiRouter.Mount(apiRouter.GetFullPath("/locations"), r)
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	locationID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	foundLocation, err := h.service.GetLocationByID(locationID)
	if utils.HandleGetByIDErrors(w, err, foundLocation, locationID, "location") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundLocation)
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	locations, err := h.service.GetAllLocations()
	if utils.HandleGetAllErrors(w, err, locations, "locations") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, locations)
}
