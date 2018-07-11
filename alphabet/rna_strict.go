package alphabet

import "strings"

// RNAStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous DNA character set.
type RNAStrict struct{}

// Contains checks that given Letter elements are in the Alphabet
func (r *RNAStrict) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(RnaStrictLetters, letter) > 0
	}
	return
}
