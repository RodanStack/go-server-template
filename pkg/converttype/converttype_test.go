package converttype_test

import (
	"go-server-template/pkg/converttype"
	"testing"
)

func TestStringToUint_NonEmptyString(t *testing.T) {
	str := "123"
	expected := uint(123)
	num, err := converttype.StringToUint(str)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if num != expected {
		t.Errorf("Expected %v, got %v", expected, num)
	}
}

func TestStringToUint_InvalidString(t *testing.T) {
	str := "abc"
	_, err := converttype.StringToUint(str)
	if err == nil {
		t.Errorf("Expected non-nil error, got nil")
	}
}

func TestStringToUint_MaxUint(t *testing.T) {
	str := "4294967295"
	expected := uint(4294967295)
	num, err := converttype.StringToUint(str)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if num != expected {
		t.Errorf("Expected %v, got %v", expected, num)
	}
}
