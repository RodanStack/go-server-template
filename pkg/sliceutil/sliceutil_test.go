package sliceutil_test

import (
	"go-server-template/pkg/sliceutil"
	"testing"
)

func TestContains_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		value    int
		expected bool
	}{
		{"ExistingValue", []int{1, 2, 3, 4, 5}, 4, true},
		{"NonExistingValue", []int{1, 2, 3, 4, 5}, 6, false},
		{"EmptySlice", []int{}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists := sliceutil.Contains(tt.slice, tt.value)
			if exists != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, exists)
			}
		})
	}
}

func TestContains_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		value    string
		expected bool
	}{
		{"ExistingValue", []string{"apple", "banana", "cherry"}, "banana", true},
		{"NonExistingValue", []string{"apple", "banana", "cherry"}, "date", false},
		{"EmptySlice", []string{}, "apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists := sliceutil.Contains(tt.slice, tt.value)
			if exists != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, exists)
			}
		})
	}
}
