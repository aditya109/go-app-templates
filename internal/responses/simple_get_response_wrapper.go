package responses

import "github.com/aditya109/go-server-template/internal/models"

// A list of items returns in the response
// swagger:response SimpleGetResponse
type SimpleGetResponseWrapper struct {
	// All items in the system
	// in: body
	Body []models.Item
}
