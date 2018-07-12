package alphabet

import "strings"

// RnaIupac is a simple struct that satisfies the Interface
// while providing complete IUPAC RNA ambiguity characters.
type RnaIupac struct{}

// Contains checks if the given letters are found
func (r *RnaIupac) Contains(letters []byte) []bool {
	found := make([]bool, len(letters))
	for idx, letter := range letters {
		found[idx] = strings.IndexByte(RnaIupacLetters, letter) > 0
	}
	return found
}

// Length is the number of letters in RnaIupac
func (r *RnaIupac) Length() int {
	return len(RnaIupacLetters)
}
