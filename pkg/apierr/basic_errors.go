package apierr

var (
	// ApiErrBadRequest is returned when the request has missing or invalid parameters.
	ApiErrBadRequest = New(400, "Bad request")

	// ApiErrUnauthorized is returned when the request requires authentication.
	ApiErrUnauthorized = New(401, "Unauthorized")

	// ApiErrForbidden is returned when the request is not allowed.
	ApiErrForbidden = New(403, "Forbidden")

	// ApiErrNotFound is returned when the requested resource is not found.
	ApiErrNotFound = New(404, "Not found")
)
