package alphabet

import "strings"

// DnaStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous DNA character set.
type DnaStrict struct{}

// Contains checks that given Letter elements are in the Alphabet
func (d *DnaStrict) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(DnaStrictLetters, letter) > 0
	}
	return
}
