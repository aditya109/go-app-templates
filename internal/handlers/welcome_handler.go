package handlers

import (
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	responseStatusCode := 200
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(responseStatusCode)
	w.Write([]byte("welcome to server ⚡ !"))
}
