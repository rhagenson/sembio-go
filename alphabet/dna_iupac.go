package alphabet

import "strings"

// DnaIupac is a simple struct that satisfies the Alphabet interface
// while providing the comlete IUPAC DNA ambiguity characters.
type DnaIupac struct{}

// Contains checks that given Letter elements are in the Alphabet
func (d *DnaIupac) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(DnaIupacLetters, letter) > 0
	}
	return
}
