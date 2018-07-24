package alphabet

import "strings"

// Alphabet is a collection of letters
type Alphabet struct {
	letters string
	width   uint
}

// New is an Alphabet generator
func New(letters string, width uint) *Alphabet {
	return &Alphabet{
		letters: letters,
		width:   width,
	}
}

// Length is numbers of letters in the Alphabet
func (a Alphabet) Length() int {
	return len(a.String()) / int(a.Width())
}

// Contains confirms whether an array of potential letters are in the Alphabet
func (a Alphabet) Contains(letters ...string) []bool {
	found := make([]bool, len(letters))
	for i, l := range letters {
		aligned := strings.Index(a.String(), l)%int(a.Width()) == 0
		present := strings.Index(a.String(), l) > -1
		found[i] = aligned && present
	}
	return found
}

// String generates a stringified copy of the Alphabet
func (a Alphabet) String() string {
	return a.letters
}

// Width is the byte width of the Alphabet
func (a Alphabet) Width() uint {
	if a.width == 0 {
		return 1
	}
	return a.width
}
