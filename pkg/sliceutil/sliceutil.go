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

// IsBetween checks if a value of any ordered type is between a minimum and maximum value (inclusive).
//
// Example usage:
//
//	func main() {
//		var value uint = 5
//		min, max := 1, 10
//		inRange := IsBetween(value, min, max)
//		fmt.Println(inRange) // Output: true
//	}
func IsBetween[
	T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64,
](value, min, max T) bool {
	return value >= min && value <= max
}
