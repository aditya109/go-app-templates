package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/aditya109/go-server-template/internal/services"
	log "github.com/aditya109/go-server-template/pkg/logwrapper"
)

var logger = log.NewLogger()

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	responseStatusCode := 200
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(responseStatusCode)
	w.Write([]byte("welcome to server ⚡ !"))
	logger.Info("/ route was hit")
}

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
}
func GetHandlerWithQueryParameters(w http.ResponseWriter, r *http.Request) {

}
