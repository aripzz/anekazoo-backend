package helpers

import (
	"github.com/gofiber/fiber/v2"
)

// StandardResponse represents the structure for the standard API response
type StandardResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendResponse is a helper function to send JSON response
func SendResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	response := StandardResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(response)
}

// SendErrorResponse is a helper function to send JSON error response
func SendErrorResponse(c *fiber.Ctx, status int, message string) error {
	response := StandardResponse{
		Status:  status,
		Message: message,
	}
	return c.Status(status).JSON(response)
}
