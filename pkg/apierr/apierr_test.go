package apierr_test

import (
	"go-server-template/pkg/apierr"
	"testing"
)

func TestNew(t *testing.T) {
	err := apierr.New(404, "Not Found")
	if err.StatusCode != 404 {
		t.Errorf("New(404, \"Not Found\").StatusCode = %v; want 404", err.StatusCode)
	}
	if err.Message != "Not Found" {
		t.Errorf("New(404, \"Not Found\").Message = %v; want \"Not Found\"", err.Message)
	}
}

func TestNewWithEmptyMessage(t *testing.T) {
	err := apierr.New(400, "")
	if err.StatusCode != 400 {
		t.Errorf("New(400, \"\").StatusCode = %v; want 400", err.StatusCode)
	}
	if err.Message != "" {
		t.Errorf("New(400, \"\").Message = %v; want \"\"", err.Message)
	}
}
