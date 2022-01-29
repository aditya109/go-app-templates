package responses

import "github.com/aditya109/go-server-template/internal/models"

// A list of items returned in the response
// swagger:response SimpleGetItemsResponse
type SimpleGetItemsResponseWrapper struct {
	// All items in the system
	// in: body
	Body []models.Item
}
