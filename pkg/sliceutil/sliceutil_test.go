package sliceutil_test

import (
	"go-server-template/pkg/sliceutil"
	"testing"
)

func TestContains_ExistingValue(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	exists := sliceutil.Contains(numbers, 4)
	if !exists {
		t.Errorf("Expected true, got %v", exists)
	}
}

func TestContains_NonExistingValue(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	exists := sliceutil.Contains(numbers, 6)
	if exists {
		t.Errorf("Expected false, got %v", exists)
	}
}

func TestContains_EmptySlice(t *testing.T) {
	var numbers []int
	exists := sliceutil.Contains(numbers, 3)
	if exists {
		t.Errorf("Expected false, got %v", exists)
	}
}

func TestIsBetween_ValueInRange(t *testing.T) {
	value := 5
	min, max := 1, 10
	inRange := sliceutil.IsBetween(value, min, max)
	if !inRange {
		t.Errorf("Expected true, got %v", inRange)
	}
}

func TestIsBetween_ValueBelowRange(t *testing.T) {
	value := 0
	min, max := 1, 10
	inRange := sliceutil.IsBetween(value, min, max)
	if inRange {
		t.Errorf("Expected false, got %v", inRange)
	}
}

func TestIsBetween_ValueAboveRange(t *testing.T) {
	value := 15
	min, max := 1, 10
	inRange := sliceutil.IsBetween(value, min, max)
	if inRange {
		t.Errorf("Expected false, got %v", inRange)
	}
}
