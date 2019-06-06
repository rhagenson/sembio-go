package alphabet

// Struct is a collection of letters
type Struct struct {
	chars map[byte]struct{}
}

// New is an Struct generator
func New(letters string) *Struct {
	set := make(map[byte]struct{}, len(letters))
	for _, b := range []byte(letters) {
		set[b] = struct{}{}
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
func (a Struct) Contains(letters ...byte) []bool {
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
	letters := make([]byte, 0, len(a.chars))
	for k := range a.chars {
		letters = append(letters, k)
	}
	return string(letters)
}
