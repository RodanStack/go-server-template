package converttype_test

import (
	"go-server-template/pkg/converttype"
	"testing"
)

func TestStringToUint(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected uint
		wantErr  bool
	}{
		{"NonEmptyString", "123", 123, false},
		{"InvalidString", "abc", 0, true},
		{"EmptyString", "", 0, false},
		{"NegativeNumberString", "-123", 0, true},
		{"OverflowUint", "18446744073709551616", 0, true}, // Value larger than max int
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num, err := converttype.StringToUint(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Expected error: %v, got: %v", tt.wantErr, err)
			}
			if num != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, num)
			}
		})
	}
}
