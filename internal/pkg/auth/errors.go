package auth

import (
	"go-server-template/pkg/apierr"
	"net/http"
)

var (
	ErrInvalidCredentials = apierr.New(http.StatusUnauthorized, "Invalid credentials")
)
