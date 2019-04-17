package alphabet

import (
	"bytes"
	"testing"
)

/*
testing.go is a collection of test helpers which should wholly
be used by people hoping to satisfy an Interface.
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
func HasExpectedLetter(a Interface, c byte) func(t *testing.T) {
	return func(t *testing.T) {
		for _, ok := range a.Contains(c) {
			if !ok {
				t.Errorf("missing expected letter: %q", c)
			}
		}
	}
}

func NotLetters(letters []byte) []byte {
	notLetters := []byte(
		"abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	)
	for _, l := range letters {
		i := bytes.IndexByte(notLetters, l)
		if i != -1 {
			notLetters = append(notLetters[:i], notLetters[i+1:]...)
		}
	}
	return notLetters
}
