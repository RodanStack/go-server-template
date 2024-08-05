package converttype

import "strconv"

// StringToUint converts a string to a uint.
// If the string is empty, it returns 0.
// If the string is not a valid number, it returns an error.
func StringToUint(str string) (uint, error) {
	if str == "" {
		return 0, nil
	}

	parsedInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return uint(parsedInt), nil
}

// UintToString converts a uint to a string.
func UintToString(num uint) string {
	return strconv.Itoa(int(num))
}
