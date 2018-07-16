package simple

import "strings"

// Rna is a simple struct that satisfies the Interface
// while providing only the unambiguous RNA character set.
type Rna struct{}

// Contains checks if the given letters are found
func (r *Rna) Contains(letters []byte) []bool {
	found := make([]bool, len(letters))
	for idx, letter := range letters {
		found[idx] = strings.IndexByte(RnaLetters, letter) > 0
	}
	return found
}

// Length is the number of letters in Rna
func (r *Rna) Length() int {
	return len(RnaLetters)
}
