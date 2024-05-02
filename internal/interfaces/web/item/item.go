package item

import (
	"encoding/json"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/item"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service *item.Service
}

func NewItemHandler(s *item.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/items/{item_id}", h.HandleGet).Methods("GET")
	router.HandleFunc("/items", h.HandleGetAll).Methods("GET")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		log.Printf("Error parsing item_id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundItem, err := h.service.GetItemByID(int32(itemID))
	if err != nil {
		log.Printf("Error getting item: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(foundItem)
	if err != nil {
		log.Printf("Error encoding item: %v", err)
		return
	}
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItems()
	if err != nil {
		log.Printf("Error getting items: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		log.Printf("Error encoding items: %v", err)
		return
	}
}
