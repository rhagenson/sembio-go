package alphabet

import (
	"bytes"
)

// Alphabet is a collection of letters
type Alphabet struct {
	letters []byte
}

// New is an Alphabet generator
func New(letters string) *Alphabet {
	return &Alphabet{
		letters: []byte(letters),
	}
}

// Length is numbers of letters in the Alphabet
func (a Alphabet) Length() int {
	return len(a.letters)
}

// Contains confirms whether an array of potential letters are in the Alphabet
func (a Alphabet) Contains(letters ...byte) []bool {
	found := make([]bool, len(letters))
	for i, l := range letters {
		found[i] = bytes.IndexByte(a.letters, l) != -1
	}
	return found
}

// String generates a stringified copy of the Alphabet
func (a Alphabet) String() string {
	return string(a.letters)
}
