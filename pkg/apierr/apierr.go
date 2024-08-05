// Package apierr provides a simple way to create an error response for an API.
// It is used to send a JSON response with an error message and status code.
package apierr

// New creates a new ApiError instance with the given status code and message.
func New(statusCode int, message string) *Error {
	return &Error{
		StatusCode: statusCode,
		Message:    message,
	}
}
