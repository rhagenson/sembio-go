package alphabet

import (
	"bytes"
	"testing"
)

/*
testing.go is a collection of test helpers which should wholly
be used by people hoping to satisfy a Interface.
*/

// IsExpectedLength is a test helper to wrap a check for
// an alphabet.Interface implementation having the correct number of letters
func IsExpectedLength(a Interface, n int) func(t *testing.T) {
	return func(t *testing.T) {
		if a.Length() != n {
			t.Errorf("Got: %d, Want: %d", a.Length(), n)
		}
	}
}

// HasExpectedLetter is a test helper to wrap a check for
// an alphabet.Interface implementation having a given letter
func HasExpectedLetter(a Interface, c string) func(t *testing.T) {
	return func(t *testing.T) {
		for _, ok := range a.Contains(c) {
			if !ok {
				t.Errorf("missing expected letter: %q", c)
			}
		}
	}
}

// TestSplitByN splits a string into n sized chunks
func TestSplitByN(s string, n int) []string {
	if len(s)%n != 0 {
		return nil
	}
	subs := make([]string, len(s)/n)

	runes := bytes.Runes([]byte(s))
	idx := 0
	for i := 0; i < len(s)+n-1; i = i + n {
		subs[idx] = string(runes[i : i+n])
		idx++
	}
	return subs
}
