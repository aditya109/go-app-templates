package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		if route.HandlerFunction != nil {
			handler = route.HandlerFunction
			handler = ConfigureHandler(handler, route)
		} else {
			handler = route.Handler
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func ConfigureHandler(inner http.Handler, route Route) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inner.ServeHTTP(w, r)
	})
}
