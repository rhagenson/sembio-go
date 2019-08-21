package utils

// BytesToStrings converts an array of bytes to
// the equaivlanet string array
func BytesToStrings(bs []byte) []string {
	a := make([]string, len(bs))
	for i, b := range bs {
		a[i] = string(b)
	}
	return a
}

// StringsToBytes converts an array of strings to
// the equivalent of the first byte from each string
func StringsToBytes(strs []string) []byte {
	a := make([]byte, len(strs))
	for i, b := range strs {
		a[i] = b[0]
	}
	return a
}
