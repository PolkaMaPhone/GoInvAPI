package interfaces

import "github.com/gorilla/mux"

type Handler interface {
	HandleRoutes(router *mux.Router)
}
