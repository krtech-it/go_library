package dto

// ErrorResponse represents an error response
// @Description Error response object returned when an error occurs
type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}
