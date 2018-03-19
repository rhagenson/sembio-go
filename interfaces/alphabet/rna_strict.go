package alphabet

var _ Alphabet = &RNAStrict{}

// RNAStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous DNA character set.
type RNAStrict struct{}

// Letters returns the unambiguous DNA character set (ATGC)
func (d *RNAStrict) Letters() []Letter {
	return []Letter{"A", "U", "G", "C"}
}

// Valid checks that a given Letter is in the Alphabet
func (d *RNAStrict) Valid(l Letter) (valid bool) {
	valid = false
	for _, c := range d.Letters() {
		if l == c {
			valid = true
		}
	}
	return
}

// Gapped is whether RNAStrict allows for gaps or not (it does not)
func (d *RNAStrict) Gapped() bool {
	return false
}

// Ambiguous is whether RNAStrict allows for ambiguity or not (it does not)
func (d *RNAStrict) Ambiguous() bool {
	return false
}
