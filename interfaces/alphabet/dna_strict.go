package alphabet

var _ Alphabet = &DNAStrict{}

// DNAStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous DNA character set.
type DNAStrict struct{}

// Letters returns the unambiguous DNA character set (ATGC)
func (d *DNAStrict) Letters() []Letter {
	return []Letter{"A", "T", "G", "C"}
}

// Valid checks that a given Letter is in the Alphabet
func (d *DNAStrict) Valid(l Letter) (valid bool) {
	valid = false
	for _, c := range d.Letters() {
		if l == c {
			valid = true
		}
	}
	return
}

// Gapped is whether DNAStrict allows for gaps or not (it does not)
func (d *DNAStrict) Gapped() bool {
	return false
}

// Ambiguous is whether DNAStrict allows for ambiguity or not (it does not)
func (d *DNAStrict) Ambiguous() bool {
	return false
}
