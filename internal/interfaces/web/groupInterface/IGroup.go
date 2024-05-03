package groupInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/groupDomain"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	service *groupDomain.Service
}

func NewGroupHandler(s *groupDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.LoggingMiddleware("INFO"))
	apiRouter.HandleFunc("/groups/{group_id}", h.HandleGet).Methods("GET")
	apiRouter.HandleFunc("/groups", h.HandleGetAll).Methods("GET")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, err := strconv.Atoi(vars["group_id"])
	if err != nil {
		middleware.ErrorLogger.Printf("Error parsing group_id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundGroup, err := h.service.GetGroupByID(int32(groupID))
	if err != nil {
		middleware.ErrorLogger.Printf("Error getting group: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundGroup)
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	groups, err := h.service.GetAllGroups()
	if err != nil {
		middleware.ErrorLogger.Printf("Error getting groups: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, groups)
}
