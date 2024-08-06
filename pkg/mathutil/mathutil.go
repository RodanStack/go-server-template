// Package mathutil provides utility functions for mathematical operations.
// It is designed to be used as a helper package for performing common math operations.
package mathutil

// IsBetween checks if a value of any ordered type is between a minimum and maximum value (inclusive).
//
// IsBetween returns true if value is greater than or equal to min and less than or equal to max.
// It supports all types that are ordered (integers, floating-point numbers, and strings).
//
// Example usage:
//
//	func main() {
//		var value uint = 5
//		min, max := 1, 10
//		inRange := IsBetween(value, min, max)
//		fmt.Println(inRange) // Output: true
//	}
func IsBetween[T Ordered](value, min, max T) bool {
	return value >= min && value <= max
}
