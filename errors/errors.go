package errors

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func (e AppError) Error() string {
	return e.Message
}

// Existing error types
func NewUnExpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

// Additional error types
func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusConflict,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

func NewMethodNotAllowedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusMethodNotAllowed,
	}
}

func NewTooManyRequestsError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusTooManyRequests,
	}
}

// Utility method to check if a specific error is of a certain type
func IsNotFoundError(err *AppError) bool {
	return err != nil && err.Code == http.StatusNotFound
}

func IsValidationError(err *AppError) bool {
	return err != nil && err.Code == http.StatusUnprocessableEntity
}

func IsBadRequestError(err *AppError) bool {
	return err != nil && err.Code == http.StatusBadRequest
}
