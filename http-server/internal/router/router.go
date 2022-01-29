package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunction
		handler = ConfigureHandler(handler)
		router.
			Methods("GET").
			Path("/").
			Name("Welcome Router").
			Handler(handler)
	}
	return router
}

func ConfigureHandler(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inner.ServeHTTP(w, r)
	})
}
