package sliceutil_test

import (
	"go-server-template/pkg/sliceutil"
	"reflect"
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

func TestIntConcat(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "multiple non-empty slices",
			input:    [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "empty slices",
			input:    [][]int{{}, {}, {}},
			expected: []int{},
		},
		{
			name:     "mixed empty and non-empty slices",
			input:    [][]int{{}, {1, 2}, {}, {3, 4}, {}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Concat(tt.input...)
			if len(result) == 0 && len(tt.expected) == 0 {
				// Both slices are empty or nil, so they are considered equal
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("IntConcat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStringConcat(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected []string
	}{
		{
			name:     "multiple non-empty slices",
			input:    [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:     "empty slices",
			input:    [][]string{{}, {}, {}},
			expected: []string{},
		},
		{
			name:     "mixed empty and non-empty slices",
			input:    [][]string{{}, {"a", "b"}, {}, {"c", "d"}, {}, {"e", "f"}},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:     "nil slices",
			input:    [][]string{nil, nil, nil},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Concat(tt.input...)
			if len(result) == 0 && len(tt.expected) == 0 {
				// Both slices are empty or nil, so they are considered equal
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("StringConcat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStructConcat(t *testing.T) {
	type TestStruct struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		input    [][]TestStruct
		expected []TestStruct
	}{
		{
			name:     "multiple non-empty slices",
			input:    [][]TestStruct{{{1, "Alice"}, {2, "Bob"}}, {{3, "Charlie"}, {4, "Diana"}}},
			expected: []TestStruct{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Diana"}},
		},
		{
			name:     "empty slices",
			input:    [][]TestStruct{{}, {}, {}},
			expected: []TestStruct{},
		},
		{
			name:     "mixed empty and non-empty slices",
			input:    [][]TestStruct{{}, {{1, "Alice"}, {2, "Bob"}}, {}, {{3, "Charlie"}, {4, "Diana"}}, {}},
			expected: []TestStruct{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Diana"}},
		},
		{
			name:     "nil slices",
			input:    [][]TestStruct{nil, nil, nil},
			expected: []TestStruct{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Concat(tt.input...)
			if len(result) == 0 && len(tt.expected) == 0 {
				// Both slices are empty or nil, so they are considered equal
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("StructConcat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPtrStructConcat(t *testing.T) {
	type TestStruct struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		input    [][]*TestStruct
		expected []*TestStruct
	}{
		{
			name: "multiple non-empty slices",
			input: [][]*TestStruct{
				{{1, "Alice"}, {2, "Bob"}},
				{{3, "Charlie"}, {4, "Diana"}},
			},
			expected: []*TestStruct{
				{1, "Alice"}, {2, "Bob"},
				{3, "Charlie"}, {4, "Diana"},
			},
		},
		{
			name:     "empty slices",
			input:    [][]*TestStruct{{}, {}, {}},
			expected: []*TestStruct{},
		},
		{
			name: "mixed empty and non-empty slices",
			input: [][]*TestStruct{
				{},
				{{1, "Alice"}, {2, "Bob"}},
				{},
				{{3, "Charlie"}, {4, "Diana"}},
				{},
			},
			expected: []*TestStruct{
				{1, "Alice"}, {2, "Bob"},
				{3, "Charlie"}, {4, "Diana"},
			},
		},
		{
			name:     "nil slices",
			input:    [][]*TestStruct{nil, nil, nil},
			expected: []*TestStruct{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Concat(tt.input...)
			if len(result) == 0 && len(tt.expected) == 0 {
				// Both slices are empty or nil, so they are considered equal
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PtrStructConcat() = %v, want %v", result, tt.expected)
			}
		})
	}
}
