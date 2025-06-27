package models

import "errors"

// Domain errors
var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrInvalidTitle  = errors.New("title must be a non-empty string")
	ErrInvalidTaskID = errors.New("invalid task ID")
)

// ErrorResponse represents error response structure
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(err string, message string) *ErrorResponse {
	return &ErrorResponse{
		Error:   err,
		Message: message,
	}
}
