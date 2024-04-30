package service

import (
	"github.com/gorilla/mux"
)

func NewRouter(handler *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HandleRequest).Methods("GET")
	r.HandleFunc("/items/{item_id}", handler.HandleGetItem).Methods("GET")
	return r
}
