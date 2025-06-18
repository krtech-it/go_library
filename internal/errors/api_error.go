package errors

type APIError struct {
	Message    string `json:"error"`
	StatusCode int    `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}
