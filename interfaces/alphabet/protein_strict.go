package alphabet

import "bitbucket.org/rhagenson/bigr/interfaces/letter"

var _ Alphabet = &ProteinStrict{}

// ProteinStrict is a simple struct that satisfies the Alphabet interface
// while providing only the unambiguous amino acid character set.
type ProteinStrict struct{}

// Letters returns the unambiguous DNA character set (ATGC)
func (p *ProteinStrict) Letters() []letter.Letter {
	return []letter.Letter{
		"A", "C", "D", "E", "F",
		"G", "H", "I", "K", "L",
		"M", "N", "P", "Q", "R",
		"S", "T", "V", "W", "Y",
	}
}

// Contains checks that given Letter elements are in the Alphabet
func (p *ProteinStrict) Contains(letter ...letter.Letter) (valid []bool) {
	for idx, letter := range letter {
		for _, inalpha := range p.Letters() {
			if letter == inalpha {
				valid[idx] = true
				continue
			}
		}
	}

	return
}

// Gapped is whether Protein allows for gaps or not (it does not)
func (p *ProteinStrict) Gapped() bool {
	return false
}

// Ambiguous is whether Protein allows for ambiguity or not (it does not)
func (p *ProteinStrict) Ambiguous() bool {
	return false
}
