// Package sliceutil provides utility functions for working with slices.
package sliceutil

// Contains checks if a slice contains a specific value.
// It returns true if the value is found, otherwise false.
//
// Example usage:
//
//	func main() {
//		numbers := []int{1, 2, 3, 4, 5}
//		exists := Contains(numbers, 3)
//		fmt.Println(exists) // Output: true
//	}
func Contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Concat concatenates multiple slices into a single slice.
func Concat[T any](slices ...[]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}
