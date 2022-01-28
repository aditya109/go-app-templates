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

func GetHandler(w http.ResponseWriter, r *http.Request) {
	responseStatusCode := 200
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(responseStatusCode)
	w.Write([]byte("get route working perfectly ⚡ !"))
}
func GetHandlerWithQueryParameters(w http.ResponseWriter, r *http.Request) {

}
