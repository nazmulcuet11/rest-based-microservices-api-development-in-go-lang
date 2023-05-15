package errs

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"error,omitempty"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func InternalServerError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error: " + message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: "Validation error: " + message,
	}
}
