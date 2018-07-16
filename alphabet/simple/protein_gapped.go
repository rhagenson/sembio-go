package simple

import "strings"

// ProteinGapped is a simple struct that satisfies the Interface
// while providing a manner to represent gaps in amino acid chains.
type ProteinGapped struct{}

// Contains checks that given Letter elements are in the Alphabet
func (p *ProteinGapped) Contains(letters []byte) []bool {
	found := make([]bool, len(letters))
	for idx, letter := range letters {
		found[idx] = strings.IndexByte(ProteinGappedLetters, letter) > 0
	}
	return found
}

// Length is the number of letters in ProteinGapped
func (p *ProteinGapped) Length() int {
	return len(ProteinGappedLetters)
}
