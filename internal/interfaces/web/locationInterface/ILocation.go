package locationInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/locationDomain"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	service *locationDomain.Service
}

func NewLocationHandler(s *locationDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.LoggingMiddleware("INFO"))
	apiRouter.HandleFunc("/locations/{location_id}", h.HandleGet).Methods("GET")
	apiRouter.HandleFunc("/locations", h.HandleGetAll).Methods("GET")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	locationID, err := strconv.Atoi(vars["location_id"])
	if err != nil {
		middleware.ErrorLogger.Printf("Error parsing location_id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundLocation, err := h.service.GetLocationByID(int32(locationID))
	if err != nil {
		middleware.ErrorLogger.Printf("Error getting location: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundLocation)
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	locations, err := h.service.GetAllLocations()
	if err != nil {
		middleware.ErrorLogger.Printf("Error getting locations: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, locations)
}
