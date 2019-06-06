package alphabet

import (
	"bytes"
	"testing"
)

/*
testing.go is a collection of test helpers which should wholly
be used by people hoping to satisfy an Interface.
*/

// TestIsExpectedLength is a test helper to wrap a check for
// an alphabet.Interface implementation having the correct number of letters
func TestIsExpectedLength(a Interface, n int) func(t *testing.T) {
	return func(t *testing.T) {
		if a.Length() != n {
			t.Errorf("Got: %d, Want: %d", a.Length(), n)
		}
	}
}

// TestHasExpectedLetter is a test helper to wrap a check for
// an alphabet.Interface implementation having a given letter
func TestHasExpectedLetter(a Interface, c byte) func(t *testing.T) {
	return func(t *testing.T) {
		for _, ok := range a.Contains(c) {
			if !ok {
				t.Errorf("missing expected letter: %q", c)
			}
		}
	}
}

// TestExcludesLetters a test helper that returns all
// ASCII letters not in the input set
func TestExcludesLetters(letters []byte) []byte {
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
