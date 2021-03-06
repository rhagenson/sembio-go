package utils

// InStrings checks if a given string is in an array.
func InStrings(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
