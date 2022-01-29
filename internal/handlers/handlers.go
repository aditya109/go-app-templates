package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	resp "github.com/aditya109/go-server-template/internal/responses"
	svc "github.com/aditya109/go-server-template/internal/services"
	log "github.com/aditya109/go-server-template/pkg/logwrapper"
)

var logger = log.NewLogger()

// WelcomeHandler returns welcome message for home URL
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	responseStatusCode := 200
	var response resp.WelcomeResponseWrapper = "welcome to server ⚡ !"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(responseStatusCode)
	w.Write([]byte(response))
	logger.Info(fmt.Sprintf("/ route was hit, status: %d", responseStatusCode))
}

// SimpleGetHandler returns a static list of items, no query params required
func SimpleGetHandler(w http.ResponseWriter, r *http.Request) {
	var responseStatusCode int
	items, err := svc.GetAllItems()
	if err != nil {
		responseStatusCode = 504
		logger.Error(err)
		w.Write([]byte(fmt.Sprintf("error occurred while getting items, %s", err.Error())))
	} else {
		responseStatusCode = 200
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
	w.WriteHeader(responseStatusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logger.Info(fmt.Sprintf("/items route was hit, status: %d", responseStatusCode))
}
func GetHandlerWithQueryParameters(w http.ResponseWriter, r *http.Request) {

}