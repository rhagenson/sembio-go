package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(AltYeast)
var _ Translater = new(AltYeast)
var _ AltNamer = new(AltYeast)
var _ IDer = new(AltYeast)
var _ StartCodoner = new(AltYeast)
var _ StopCodoner = new(AltYeast)

type (
	// AltYeast is the alternative yeast DNA to protein translation table
	AltYeast struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s AltYeast) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGG": 'W', "CTT": 'L', "CTC": 'L', "CTA": 'L',
		"CTG": 'S', "CCT": 'P', "CCC": 'P', "CCA": 'P',
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
		"GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s AltYeast) String() string {
	return "Alternative Yeast Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s AltYeast) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s AltYeast) ID() uint {
	return 12
}

// StartCodons lists the codons which start a transcript
func (s AltYeast) StartCodons() []string {
	return []string{"CTG", "ATG"}
}

// StopCodons lists the codons which end a transcript
func (s AltYeast) StopCodons() []string {
	return []string{"TAA", "TAG", "TGA"}
}
