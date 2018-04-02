package alphabet

import "bitbucket.org/rhagenson/bigr/interfaces/letter"

var _ Alphabet = &RNAIupac{}

// RNAIupac is a simple struct that satisfies the Alphabet interface
// while providing the comlete IUPAC DNA ambiguity characters.
type RNAIupac struct{}

// Letters returns the complete IUPAC ambiguity character set.
func (r *RNAIupac) Letters() []letter.Letter {
	return []letter.Letter{
		"A", "U", "G", "C", // Any of one nucelotide codes (i.e., 4 choose 1)
		"R", "Y", "S", "W", "K", "M", // Any of two nucelotide codes (i.e., 4 choose 2)
		"B", "D", "H", "V", "N", // Any of three nucleotide codes (i.e., 4 choose 3)
		"-", // Gap code (i.e., 4 choose 0)
	}
}

// Contains checks that given Letter elements are in the Alphabet
func (r *RNAIupac) Contains(letter ...letter.Letter) (valid []bool) {
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

// Gapped is whether RNAIupac allows for gaps or not (it does)
func (r *RNAIupac) Gapped() bool {
	return true
}

// Ambiguous is whether RNAIupac allows for ambiguous characters or not
// (it does)
func (r *RNAIupac) Ambiguous() bool {
	return true
}
