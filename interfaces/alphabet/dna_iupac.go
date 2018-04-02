package alphabet

import "bitbucket.org/rhagenson/bigr/interfaces/letter"

var _ Alphabet = &DNAIupac{}

// DNAIupac is a simple struct that satisfies the Alphabet interface
// while providing the comlete IUPAC DNA ambiguity characters.
type DNAIupac struct{}

// Letters returns the complete IUPAC ambiguity character set.
func (d *DNAIupac) Letters() []letter.Letter {
	return []letter.Letter{
		"A", "T", "G", "C", // Any of one nucelotide codes (i.e., 4 choose 1)
		"R", "Y", "S", "W", "K", "M", // Any of two nucelotide codes (i.e., 4 choose 2)
		"B", "D", "H", "V", "N", // Any of three nucleotide codes (i.e., 4 choose 3)
		"-", // Gap code (i.e., 4 choose 0)
	}
}

// Contains checks that given Letter elements are in the Alphabet
func (d *DNAIupac) Contains(letter ...letter.Letter) (valid []bool) {
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

// Gapped is whether DNAIupac allows for gaps or not (it does)
func (d *DNAIupac) Gapped() bool {
	return true
}

// Ambiguous is whether DNAIupac allows for ambiguous characters or not
// (it does)
func (d *DNAIupac) Ambiguous() bool {
	return true
}
