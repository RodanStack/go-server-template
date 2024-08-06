package apierr_test

import (
	"go-server-template/pkg/apierr"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		message    string
		wantStatus int
		wantMsg    string
	}{
		{"NotFound", http.StatusNotFound, "Not Found", http.StatusNotFound, "Not Found"},
		{"BadRequest", http.StatusBadRequest, "Bad Request", http.StatusBadRequest, "Bad Request"},
		{
			"InternalServerError",
			http.StatusInternalServerError,
			"Internal Server Error",
			http.StatusInternalServerError,
			"Internal Server Error",
		},
		{"EmptyMessage", http.StatusBadRequest, "", http.StatusBadRequest, ""},
		{"ZeroStatusCode", 0, "Zero Status Code", 0, "Zero Status Code"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := apierr.New(tt.statusCode, tt.message)
			if err.StatusCode != tt.wantStatus {
				t.Errorf("New(%d, %q).StatusCode = %v; want %v", tt.statusCode, tt.message, err.StatusCode, tt.wantStatus)
			}
			if err.Message != tt.wantMsg {
				t.Errorf("New(%d, %q).Message = %v; want %v", tt.statusCode, tt.message, err.Message, tt.wantMsg)
			}
		})
	}
}
