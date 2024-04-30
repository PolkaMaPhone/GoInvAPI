package handlers

import "net/http"

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Handle your requests here
	w.Write([]byte("Hello, World!"))
}
