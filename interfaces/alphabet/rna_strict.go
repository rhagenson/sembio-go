package alphabet

var _ Alphabet = &RNAStrict{}

// RNAStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous DNA character set.
type RNAStrict struct{}

// Letters returns the unambiguous DNA character set (ATGC)
func (r *RNAStrict) Letters() []Letter {
	return []Letter{"A", "U", "G", "C"}
}

// Contains checks that given Letter elements are in the Alphabet
func (r *RNAStrict) Contains(letter ...Letter) (valid []bool) {
	for idx, letter := range letter {
		for _, inalpha := range r.Letters() {
			if letter == inalpha {
				valid[idx] = true
				continue
			}
		}
	}

	return
}

// Gapped is whether RNAStrict allows for gaps or not (it does not)
func (r *RNAStrict) Gapped() bool {
	return false
}

// Ambiguous is whether RNAStrict allows for ambiguity or not (it does not)
func (r *RNAStrict) Ambiguous() bool {
	return false
}
