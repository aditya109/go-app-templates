package router

import (
	"net/http"

	h "github.com/aditya109/go-server-template/internal/handlers"
)

type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunction http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"WelcomeHandler",
		"GET",
		"/",
		h.WelcomeHandler,
	},
	Route{
		"SimpleGetHandler",
		"GET",
		"/items",
		h.SimpleGetHandler,
	},
}
