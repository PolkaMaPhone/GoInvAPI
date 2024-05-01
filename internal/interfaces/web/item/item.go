package item

import (
	"context"
	"encoding/json"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	DB db.DBTX
}

func NewItemHandler(db db.DBTX) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/items/{item_id}", h.HandleGet).Methods("GET")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		log.Printf("Error parsing item_id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries := db.New(h.DB)
	item, err := queries.GetItem(context.Background(), int32(itemID))
	if err != nil {
		log.Printf("Error getting item: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Printf("Error encoding item: %v", err)
		return
	}
}
