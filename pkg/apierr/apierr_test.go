package apierr_test

import (
	"go-server-template/pkg/apierr"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	err := apierr.New(http.StatusNotFound, "Not Found")
	if err.StatusCode != http.StatusNotFound {
		t.Errorf("New(404, \"Not Found\").StatusCode = %v; want 404", err.StatusCode)
	}
	if err.Message != "Not Found" {
		t.Errorf("New(404, \"Not Found\").Message = %v; want \"Not Found\"", err.Message)
	}
}

func TestNewWithEmptyMessage(t *testing.T) {
	err := apierr.New(http.StatusBadRequest, "")
	if err.StatusCode != http.StatusBadRequest {
		t.Errorf("New(400, \"\").StatusCode = %v; want 400", err.StatusCode)
	}
	if err.Message != "" {
		t.Errorf("New(400, \"\").Message = %v; want \"\"", err.Message)
	}
}
