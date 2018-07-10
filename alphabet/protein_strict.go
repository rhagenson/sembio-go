package alphabet

import "strings"

// ProteinStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous amino acid character set.
type ProteinStrict struct{}

// Contains checks that given Letter elements are in the Alphabet
func (p *ProteinStrict) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(dnaIupacLetters, letter) > 0
	}
	return
}
