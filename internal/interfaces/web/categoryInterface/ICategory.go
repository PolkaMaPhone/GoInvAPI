package categoryInterface

import (
	"encoding/json"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/categoryDomain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service *categoryDomain.Service
}

func NewCategoryHandler(s *categoryDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/categories/{category_id}", h.HandleGet).Methods("GET")
	apiRouter.HandleFunc("/categories", h.HandleGetAll).Methods("GET")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		log.Printf("Error parsing category_id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundCategory, err := h.service.GetCategoryByID(int32(categoryID))
	if err != nil {
		log.Printf("Error getting category: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(foundCategory)
	if err != nil {
		log.Printf("Error encoding category: %v", err)
		return
	}
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		log.Printf("Error getting categories: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		log.Printf("Error encoding categories: %v", err)
		return
	}
}
