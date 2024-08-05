package ptrutil_test

import (
	"go-server-template/pkg/ptrutil"
	"testing"
)

func TestToPointer(t *testing.T) {
	var num uint = 42
	ptr := ptrutil.ToPointer(num)
	if *ptr != 42 {
		t.Errorf("ToPointer(%v) = %v; want 42", num, *ptr)
	}
}

func TestToValueWithNilPointer(t *testing.T) {
	var num *uint
	value := ptrutil.ToValue(num)
	if value != 0 {
		t.Errorf("ToValue(%v) = %v; want 0", num, value)
	}
}

func TestToPointerWithNilValue(t *testing.T) {
	var num *uint
	ptr := ptrutil.ToPointer(num)
	if ptr != nil {
		t.Errorf("ToPointer(%v) = %v; want nil", num, ptr)
	}
}

func TestToValueWithStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "John Doe", Age: 30}
	ptr := &person
	value := ptrutil.ToValue(ptr)
	if value != person {
		t.Errorf("ToValue(%v) = %v; want %v", ptr, value, person)
	}
}

func TestToPointerWithStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "John Doe", Age: 30}
	ptr := ptrutil.ToPointer(person)
	if *ptr != person {
		t.Errorf("ToPointer(%v) = %v; want %v", person, *ptr, person)
	}
}
