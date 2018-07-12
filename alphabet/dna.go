package alphabet

import "strings"

// Dna is a simple struct that satisfies the Interface
// while providing only the unambiguous DNA character set.
type Dna struct{}

// Contains checks if the given letters are found
func (d *Dna) Contains(letters []byte) []bool {
	found := make([]bool, len(letters))
	for idx, letter := range letters {
		found[idx] = strings.IndexByte(DnaLetters, letter) > 0
	}
	return found
}

// Length is the number of letters in Dna
func (d *Dna) Length() int {
	return len(DnaLetters)
}
