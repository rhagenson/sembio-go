package alphabet

import "strings"

// DNAStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous DNA character set.
type DNAStrict struct{}

// Contains checks that given Letter elements are in the Alphabet
func (d *DNAStrict) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(dnaStrictLetters, letter) > 0
	}
	return
}
