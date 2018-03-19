package alphabet

var _ Alphabet = &ProteinGapped{}

// ProteinGapped is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous amino acid character set.
type ProteinGapped struct{}

// Letters returns the unambiguous DNA character set (ATGC)
func (p *ProteinGapped) Letters() []Letter {
	return []Letter{
		"A", "C", "D", "E", "F",
		"G", "H", "I", "K", "L",
		"M", "N", "P", "Q", "R",
		"S", "T", "V", "W", "Y",
		"-",
	}
}

// Valid checks that a given Letter is in the Alphabet
func (p *ProteinGapped) Valid(l Letter) (valid bool) {
	valid = false
	for _, c := range p.Letters() {
		if l == c {
			valid = true
		}
	}
	return
}

// Gapped is whether Protein allows for gaps or not (it does not)
func (p *ProteinGapped) Gapped() bool {
	return true
}

// Ambiguous is whether Protein allows for ambiguity or not (it does not)
func (p *ProteinGapped) Ambiguous() bool {
	return false
}
