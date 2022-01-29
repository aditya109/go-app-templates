package router

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunction
		handler = ConfigureHandler(handler, route)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// configure swagger route
	opts := middleware.RedocOpts{
		SpecURL: "/swagger.yaml",
	}
	sh := middleware.Redoc(opts, nil)
	router.Handle("/docs", sh)
	router.
		Methods("GET").
		Path("/swagger.yaml").
		Name("Swagger JSON").
		Handler(http.FileServer(http.Dir("./api/swagger")))
	return router
}

func ConfigureHandler(inner http.Handler, route Route) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inner.ServeHTTP(w, r)
	})
}
