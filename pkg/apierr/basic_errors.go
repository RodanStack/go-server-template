package apierr

import "net/http"

var (
	// ErrBadRequest is returned when the request has missing or invalid parameters.
	ErrBadRequest = New(http.StatusBadRequest, "Bad request")

	// ErrUnauthorized is returned when the request requires authentication.
	ErrUnauthorized = New(http.StatusUnauthorized, "Unauthorized")

	// ErrForbidden is returned when the request is not allowed.
	ErrForbidden = New(http.StatusForbidden, "Forbidden")

	// ErrNotFound is returned when the requested resource is not found.
	ErrNotFound = New(http.StatusNotFound, "Not found")
)
