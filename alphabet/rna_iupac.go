package alphabet

import "strings"

// RNAIupac is a simple struct that satisfies the Alphabet interface
// while providing the comlete IUPAC DNA ambiguity characters.
type RNAIupac struct{}

// Contains checks that given Letter elements are in the Alphabet
func (r *RNAIupac) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(rnaIupacLetters, letter) > 0
	}
	return
}
