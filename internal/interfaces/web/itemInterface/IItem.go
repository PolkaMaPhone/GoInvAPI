package itemInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	service *itemDomain.Service
}

func NewItemHandler(s *itemDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.LoggingMiddleware("INFO"))
	apiRouter.HandleFunc("/items/{item_id}", h.HandleGet).Methods("GET")
	apiRouter.HandleFunc("/items", h.HandleGetAll).Methods("GET")
	apiRouter.HandleFunc("/items/{item_id}/with_category", h.HandleGetWithCategory).Methods("GET")
	apiRouter.HandleFunc("/items_with_category", h.HandleGetAllWithCategories).Methods("GET")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		middleware.ErrorLogger.Printf("Error parsing item_id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundItem, err := h.service.GetItemByID(int32(itemID))
	if err != nil {
		middleware.ErrorLogger.Printf("Error getting item: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItems()
	if err != nil {
		middleware.ErrorLogger.Printf("Error getting items: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}

func (h *Handler) HandleGetWithCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		middleware.ErrorLogger.Printf("Error parsing item_id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundItem, err := h.service.GetItemByIDWithCategory(int32(itemID))
	if err != nil {
		middleware.ErrorLogger.Printf("Error getting item with category: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetAllWithCategories(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithCategories()
	if err != nil {
		middleware.ErrorLogger.Printf("Error getting items with category: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}
