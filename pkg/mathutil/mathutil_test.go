package mathutil_test

import (
	"go-server-template/pkg/mathutil"
	"testing"
)

func TestIsBetween_Int(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		min      int
		max      int
		expected bool
	}{
		{"ValueEqualToMin", 1, 1, 10, true},
		{"ValueEqualToMax", 10, 1, 10, true},
		{"ValueWithinRange", 5, 1, 10, true},
		{"ValueBelowMin", 0, 1, 10, false},
		{"ValueAboveMax", 11, 1, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inRange := mathutil.IsBetween(tt.value, tt.min, tt.max)
			if inRange != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, inRange)
			}
		})
	}
}

func TestIsBetween_NegativeInt(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		min      int
		max      int
		expected bool
	}{
		{"ValueEqualToMin", -10, -10, -1, true},
		{"ValueEqualToMax", -1, -10, -1, true},
		{"ValueWithinRange", -5, -10, -1, true},
		{"ValueBelowMin", -11, -10, -1, false},
		{"ValueAboveMax", 0, -10, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inRange := mathutil.IsBetween(tt.value, tt.min, tt.max)
			if inRange != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, inRange)
			}
		})
	}
}

func TestIsBetween_String(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		min      string
		max      string
		expected bool
	}{
		{"ValueEqualToMin", "apple", "apple", "banana", true},
		{"ValueEqualToMax", "banana", "apple", "banana", true},
		{"ValueWithinRange", "cat", "apple", "dog", true},
		{"ValueBelowMin", "aardvark", "apple", "banana", false},
		{"ValueAboveMax", "orange", "apple", "banana", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inRange := mathutil.IsBetween(tt.value, tt.min, tt.max)
			if inRange != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, inRange)
			}
		})
	}
}
