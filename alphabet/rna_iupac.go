package alphabet

import "strings"

// RnaIupac is a simple struct that satisfies the Alphabet interface
// while providing the comlete IUPAC DNA ambiguity characters.
type RnaIupac struct{}

// Contains checks that given Letter elements are in the Alphabet
func (r *RnaIupac) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(RnaIupacLetters, letter) > 0
	}
	return
}
