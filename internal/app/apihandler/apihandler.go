package apihandler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/db"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type APIHandler struct {
	DB db.DBTX
}

func NewAPIHandler() *APIHandler {
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	DbSchema := os.Getenv("DB_SCHEMA")
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", DbUser, DbPassword, DbHost, DbPort, DbName, DbSchema)

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return &APIHandler{
		DB: conn,
	}
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
