package handlers

import "net/http"

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Handle your requests here
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		return
	}
}
