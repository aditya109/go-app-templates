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
		"Welcome",
		"GET",
		"/",
		h.WelcomeHandler,
	},
	// swagger:route GET /items items listItems
	// Returns a list of items, no query params required
	// responses:
	// 	200: SimpleGetItemsResponse
	Route{
		"SimpleGetItems",
		"GET",
		"/items",
		h.SimpleGetItemsHandler,
	},
	// swagger:route GET /item/{id} item listItem
	// Returns an item with id from the existing list of items
	// responses:
	// 	200: GetItemWithIdReponse
	Route{
		"GetItemWithId",
		"GET",
		"/item/{id}",
		h.GetItemWithIdHandler,
	},

	// TODO:
	// swagger:route GET /items?id=12&name=Bertine items listItems
	// Returns items filtered by query parameters from the existing list of items
	// responses:
	// 	200: GetWithQueryParamsReponse
	Route{
		"GetWithQueryParams",
		"GET",
		"/item/{id}",
		h.GetWithQueryParamsHandler,
	},
	// TODO:
	// swagger:route GET /item/{id} item listItem
	// Returns an item with id from the existing list of items
	// responses:
	// 	200: GetItemWithIdReponse
	Route{
		"GetItemWithIdHandler",
		"GET",
		"/item/{id}",
		h.GetItemWithIdHandler,
	},
}
