package alphabet

var _ Alphabet = &DNAStrict{}

// DNAStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous DNA character set.
type DNAStrict struct{}

// Letters returns the unambiguous DNA character set (ATGC)
func (d *DNAStrict) Letters() []Letter {
	return []Letter{"A", "T", "G", "C"}
}

// Contains checks that given Letter elements are in the Alphabet
func (d *DNAStrict) Contains(letter ...Letter) (valid []bool) {
	for idx, letter := range letter {
		for _, inalpha := range d.Letters() {
			if letter == inalpha {
				valid[idx] = true
				continue
			}
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
