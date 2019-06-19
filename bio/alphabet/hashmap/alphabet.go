package hashmap

import "strings"

// Struct is a collection of letters
type Struct struct {
	chars map[string]struct{}
}

// New is an Struct generator
func New(letters string) *Struct {
	set := make(map[string]struct{}, len(letters))
	for _, b := range letters {
		set[string(b)] = struct{}{}
	}
	return &Struct{
		chars: set,
	}
}

// Length is numbers of letters in the Alphabet
func (a Struct) Length() int {
	return len(a.chars)
}

// Contains confirms whether a potential letters are in the Alphabet
func (a Struct) Contains(letters ...string) []bool {
	found := make([]bool, len(letters))
	for i, l := range letters {
		if _, ok := a.chars[l]; ok {
			found[i] = true
		}
	}
	return found
}

// String generates a stringified copy of the Alphabet
func (a Struct) String() string {
	letters := make([]string, 0, len(a.chars))
	for k := range a.chars {
		letters = append(letters, k)
	}
	return strings.Join(letters, "")
}
