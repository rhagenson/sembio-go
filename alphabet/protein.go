package alphabet

import "strings"

// Protein is a simple struct that satisfies the Interface
// while providing only the unambiguous amino acid character set.
type Protein struct{}

// Contains checks if the given letters are found
func (p *Protein) Contains(letters []byte) []bool {
	found := make([]bool, len(letters))
	for idx, letter := range letters {
		found[idx] = strings.IndexByte(ProteinLetters, letter) > 0
	}
	return found
}

// Length is the number of letters in Protein
func (p *Protein) Length() int {
	return len(ProteinLetters)
}
