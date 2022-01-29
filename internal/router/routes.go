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

	// swagger:route GET / home welcome
	// Returns a welcome message
	// responses:
	// 	200: WelcomeResponse
	Route{
		"WelcomeHandler",
		"GET",
		"/",
		h.WelcomeHandler,
	},
	// swagger:route GET /items items listItems
	// Returns a list of items, no query params required
	// responses:
	// 	200: SimpleGetResponse
	Route{
		"SimpleGetHandler",
		"GET",
		"/items",
		h.SimpleGetHandler,
	},
}
