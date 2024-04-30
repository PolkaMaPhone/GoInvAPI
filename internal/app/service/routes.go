package service

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/app/apihandler"
	"github.com/gorilla/mux"
)

func NewRouter(apiHandler *apihandler.APIHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", apiHandler.HandleRequest).Methods("GET")
	r.HandleFunc("/items/{item_id}", apiHandler.HandleGetItem).Methods("GET")
	return r
}
