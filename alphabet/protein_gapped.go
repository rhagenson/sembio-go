package alphabet

import "strings"

// ProteinGapped is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous amino acid character set.
type ProteinGapped struct{}

// Contains checks that given Letter elements are in the Alphabet
func (p *ProteinGapped) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(DnaIupacLetters, letter) > 0
	}
	return
}
