package utils

import "strings"

// MakeExcited transforms a setence to all caps with an exclamation point
func MakeExcited(sentence string) string {
	return strings.ToUpper(sentence) + "!"
}
