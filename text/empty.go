package text

import "strings"

func IsEmpty(str string) bool {
	if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}

// IsNotEmpty checks if the string is not empty.
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}
