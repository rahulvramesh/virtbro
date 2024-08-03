package utils

import (
	"strings"
)

// CapitalizeString capitalizes the first letter of the given string
func CapitalizeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
