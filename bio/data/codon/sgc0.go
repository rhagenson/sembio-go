package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC0)
var _ Translater = new(SGC0)
var _ AltNamer = new(SGC0)
var _ IDer = new(SGC0)
var _ StartCodoner = new(SGC0)
var _ StopCodoner = new(SGC0)

type (
	// SGC0 is the NCBI DNA to Protein translation table
	SGC0 struct{}

	// Standard is the standard DNA to Protein translation table
	Standard struct {
		SGC0
	}
)

// Translate converts a codon into its amino acid equivalent
func (s SGC0) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGG": 'W', "CTT": 'L', "CTC": 'L', "CTA": 'L',
		"CTG": 'L', "CCT": 'P', "CCC": 'P', "CCA": 'P',
		"CCG": 'P', "CAT": 'H', "CAC": 'H', "CAA": 'Q',
		"CAG": 'Q', "CGT": 'R', "CGC": 'R', "CGA": 'R',
		"CGG": 'R', "ATT": 'I', "ATC": 'I', "ATA": 'I',
		"ATG": 'M', "ACT": 'T', "ACC": 'T', "ACA": 'T',
		"ACG": 'T', "AAT": 'N', "AAC": 'N', "AAA": 'K',
		"AAG": 'K', "AGT": 'S', "AGC": 'S', "AGA": 'R',
		"AGG": 'R', "GTT": 'V', "GTC": 'V', "GTA": 'V',
		"GTG": 'V', "GCT": 'A', "GCC": 'A', "GCA": 'A',
		"GCG": 'A', "GAT": 'D', "GAC": 'D', "GAA": 'E',
		"GAG": 'E', "GGT": 'G', "GGC": 'G', "GGA": 'G',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s SGC0) String() string {
	return "SGC0 Codon Library"
}

// String provides a human-readable indication of usage
func (s Standard) String() string {
	return "Standard Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s SGC0) AltName() string {
	return "SGC0"
}

// ID provides the identifier used by NCBI
func (s SGC0) ID() uint {
	return 1
}

// StartCodons lists the codons which start a transcript
func (s SGC0) StartCodons() []string {
	return []string{"TTG", "CTG", "ATG"}
}

// StopCodons lists the codons which end a transcript
func (s SGC0) StopCodons() []string {
	return []string{"TAA", "TAG", "TGA"}
}
