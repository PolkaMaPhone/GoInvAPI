package locationInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/locationDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validationMiddleware"
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
	apiRouter.Route("/api/locations", func(r chi.Router) {
		r.Route("/{location_id}", func(r chi.Router) {
			r.Use(validationMiddleware.ValidateMethod(http.MethodGet, http.MethodPut, http.MethodDelete))
			r.Get("/", h.HandleGet)
			r.Delete("/", h.HandleDelete)
			r.Put("/", h.HandlePut)
		})

		r.With(validationMiddleware.ValidateMethod(http.MethodPost)).Post("/", h.HandlePost)

		r.With(validationMiddleware.ValidateMethod(http.MethodGet)).Get("/", h.HandleGetAll)
	})
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	locationID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	logging.InfoLogger.Printf("Retrieved locationID: %v", locationID)

	foundLocation, err := h.service.GetLocationByID(locationID)
	if err != nil {
		utils.HandleGetByIDErrors(w, err, foundLocation, locationID, "location")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, foundLocation)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	locations, err := h.service.GetAllLocations()
	if err != nil {
		utils.HandleGetAllErrors(w, err, locations, "locations")
		return
	}
	logging.InfoLogger.Printf("Retrieved all locations: %v", locations)

	err = utils.RespondWithJSON(w, http.StatusOK, locations)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	logging.ErrorLogger.Printf("Location Delete not implemented yet")
	http.Error(writer, "Location Delete not implemented yet", http.StatusNotImplemented)
	//TODO - Implement Location Delete
}

func (h *Handler) HandlePut(writer http.ResponseWriter, request *http.Request) {
	logging.ErrorLogger.Printf("Location update not implemented yet")
	http.Error(writer, "Location update not implemented yet", http.StatusNotImplemented)
	//TODO - Implement Location update
}

func (h *Handler) HandlePost(writer http.ResponseWriter, request *http.Request) {
	logging.ErrorLogger.Printf("Location creation not implemented yet")
	http.Error(writer, "Location creation not implemented yet", http.StatusNotImplemented)
	//TODO - Implement Location creation
}
