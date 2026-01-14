package cli

import (
	"fmt"
	"os"
)

// ExitWithError prints an error message and exits with code 1
func ExitWithError(err error) {
	Error(err.Error())
	os.Exit(1)
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// UserAbortError represents when user cancels an operation
type UserAbortError struct{}

func (e *UserAbortError) Error() string {
	return "operation cancelled by user"
}

// IsUserAbort checks if an error is a user abort
func IsUserAbort(err error) bool {
	_, ok := err.(*UserAbortError)
	return ok
}
