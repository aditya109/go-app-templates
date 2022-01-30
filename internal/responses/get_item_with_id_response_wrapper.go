package responses

import "github.com/aditya109/go-server-template/internal/models"

// GetItemWithIDResponseWrapper is an item returned in response
// swagger:response GetItemWithIdResponse
type GetItemWithIDResponseWrapper struct {
	// Item in the system
	// in: body
	Body models.Item
}
