package alphabet

var _ Alphabet = &RNAIupac{}

// RNAIupac is a simple struct that satisfies the Alphabet interface
// while providing the comlete IUPAC DNA ambiguity characters.
type RNAIupac struct{}

// Letters returns the complete IUPAC ambiguity character set.
func (d *RNAIupac) Letters() []Letter {
	return []Letter{
		"A", "U", "G", "C", // Any of one nucelotide codes (i.e., 4 choose 1)
		"R", "Y", "S", "W", "K", "M", // Any of two nucelotide codes (i.e., 4 choose 2)
		"B", "D", "H", "V", "N", // Any of three nucleotide codes (i.e., 4 choose 3)
		"-", // Gap code (i.e., 4 choose 0)
	}
}

// Valid checks that a given Letter is in the Alphabet
func (d *RNAIupac) Valid(l Letter) (valid bool) {
	valid = false
	for _, c := range d.Letters() {
		if l == c {
			valid = true
		}
	}
	return
}

// Gapped is whether RNAIupac allows for gaps or not (it does)
func (d *RNAIupac) Gapped() bool {
	return true
}

// Ambiguous is whether RNAIupac allows for ambiguous characters or not
// (it does)
func (d *RNAIupac) Ambiguous() bool {
	return true
}
