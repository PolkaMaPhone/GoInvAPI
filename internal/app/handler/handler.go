package handlers

import (
	"context"
	"encoding/json"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/db"
	"net/http"
	"strconv"
)

func HandleGetItem(w http.ResponseWriter, r *http.Request) {
	itemID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := db.GetItem(context.Background(), int32(itemID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		return
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(map[string]string{"status": "server running"})
	_, err := w.Write(response)
	if err != nil {
		return
	}
}
