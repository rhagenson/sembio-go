package utils

func BytesToStrings(bs []byte) []string {
	a := make([]string, len(bs))
	for i, b := range bs {
		a[i] = string(b)
	}
	return a
}

func StringsToBytes(strs []string) []byte {
	a := make([]byte, len(strs))
	for i, b := range strs {
		a[i] = b[0]
	}
	return a
}
