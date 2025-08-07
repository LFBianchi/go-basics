package mystrings

// Reverse a string left to right.
func Reverse(s string) string {
	result := ""
	for _, character := range s {
		result = string(character) + result
	}
	return result
}
