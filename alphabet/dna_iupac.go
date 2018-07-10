package alphabet

import "strings"

// DNAIupac is a simple struct that satisfies the Alphabet interface
// while providing the comlete IUPAC DNA ambiguity characters.
type DNAIupac struct{}

// Contains checks that given Letter elements are in the Alphabet
func (d *DNAIupac) Contains(letters []byte) (valid []bool) {
	for idx, letter := range letters {
		valid[idx] = strings.IndexByte(dnaIupacLetters, letter) > 0
	}
	return
}
