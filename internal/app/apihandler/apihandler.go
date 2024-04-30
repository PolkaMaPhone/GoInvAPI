package apihandler

import (
	"context"
	"encoding/json"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type APIHandler struct {
	DB db.DBTX
}

func (h *APIHandler) HandleGetItem(w http.ResponseWriter, r *http.Request) {
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

func (h *APIHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	port := r.Host[strings.LastIndex(r.Host, ":")+1:]
	response, _ := json.Marshal(map[string]string{"status": "server running", "port": port})
	_, err := w.Write(response)
	if err != nil {
		return
	}
}
