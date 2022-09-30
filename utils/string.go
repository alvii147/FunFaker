package utils

import "strings"

func IsStringAlphabeticallyLess(a string, b string) bool {
	lowerComparison := strings.Compare(
		strings.ToLower(a),
		strings.ToLower(b),
	)

	if lowerComparison != 0 {
		return lowerComparison == -1
	}

	return strings.Compare(
		a,
		b,
	) == -1
}

func StringSoftEqual(expected string, actual string) bool {
	if expected != "" && !strings.EqualFold(expected, actual) {
		return false
	}

	return true
}
