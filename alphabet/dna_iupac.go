package alphabet

import "strings"

// DnaIupac is a simple struct that satisfies the Interface
// while providing the comlete IUPAC DNA ambiguity characters.
type DnaIupac struct{}

// Contains checks if the given letters are found
func (d *DnaIupac) Contains(letters []byte) []bool {
	found := make([]bool, len(letters))
	for idx, letter := range letters {
		found[idx] = strings.IndexByte(DnaIupacLetters, letter) > 0
	}
	return found
}

// Length is the number of letters in DnaIupac
func (d *DnaIupac) Length() int {
	return len(DnaIupacLetters)
}
