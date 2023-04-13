package http

// ErrorResponse represents the response payload of an app error.
type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}
