package utils

// InStrings checks if a given strings is in an array.
func InStrings(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
