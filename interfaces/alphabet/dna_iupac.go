package alphabet

var _ Alphabet = &DNAIupac{}

// DNAIupac is a simple struct that satisfies the Alphabet interface
// while providing the comlete IUPAC DNA ambiguity characters.
type DNAIupac struct{}

// Letters returns the complete IUPAC ambiguity character set.
func (d *DNAIupac) Letters() []Letter {
	return []Letter{
		"A", "T", "G", "C", // Any of one nucelotide codes (i.e., 4 choose 1)
		"R", "Y", "S", "W", "K", "M", // Any of two nucelotide codes (i.e., 4 choose 2)
		"B", "D", "H", "V", "N", // Any of three nucleotide codes (i.e., 4 choose 3)
		"-", // Gap code (i.e., 4 choose 0)
	}
}

// Valid checks that a given Letter is in the Alphabet
func (d *DNAIupac) Valid(l Letter) (valid bool) {
	valid = false
	for _, c := range d.Letters() {
		if l == c {
			valid = true
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
