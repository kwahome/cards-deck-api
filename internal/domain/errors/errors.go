package errors

import "fmt"

// ServiceError represents a domain service error.
type ServiceError struct {
	Code    string
	Message string
	cause   error
}

// NewServiceError creates a new domain service error with a code and Message.
//
// param: `code`	- the error code
//
// param: `Message`	- the error Message
//
// returns: a domain service error.
func NewServiceError(code string, message string) error {
	return &ServiceError{
		Code:    code,
		Message: message,
	}
}

func (error *ServiceError) Error() string {
	if error.cause != nil {
		return fmt.Sprintf("error code: %s, Message: %s, cause: %v", error.Code, error.Message, error.cause)
	}
	return fmt.Sprintf("error code: %s, Message: %s", error.Code, error.Message)
}
