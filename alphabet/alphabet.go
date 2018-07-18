package alphabet

import "strings"

// Alphabet is a collection of letters
type Alphabet string

// Length is numbers of letters in the Alphabet
func (a Alphabet) Length() int {
	return len(a)
}

// Contains confirms whether an array of potential letters are in the Alphabet
func (a Alphabet) Contains(ls []byte) []bool {
	found := make([]bool, len(ls))
	for idx, letter := range ls {
		found[idx] = strings.IndexByte(string(a), letter) > -1
	}
	return found
}

// Copy generates a stringified copy of the Alphabet
func (a Alphabet) Copy() string {
	return string(a)
}
