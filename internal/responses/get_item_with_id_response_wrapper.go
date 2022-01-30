package responses

import "github.com/aditya109/go-server-template/internal/models"

// GetItemWithIdResponseWrapper is an item returned in response
// swagger:response GetItemWithIdResponse
type GetItemWithIdResponseWrapper struct {
	// Item in the system
	// in: body
	Body models.Item
}
