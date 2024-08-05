package pointer

// ToValue converts a pointer to its value.
// If the pointer is nil, it returns the zero value of the type.
//
// Example usage:
//
//	func main() {
//		var num uint = 42
//		ptr := &num
//		value := ToValue(ptr)
//		fmt.Println(value) // Output: 42
//	}
func ToValue[T any](ptr *T) T {
	if ptr == nil {
		var zeroValue T
		return zeroValue
	}
	return *ptr
}

// ToPointer converts any type to a pointer.
//
// Example usage:
//
//	func main() {
//		var num uint = 42
//		ptr := ToPointer(num)
//		fmt.Println(*ptr) // Output: 42
//	}
func ToPointer[T any](value T) *T {
	return &value
}
