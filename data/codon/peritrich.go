package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(Peritrich)
var _ Translater = new(Peritrich)
var _ AltNamer = new(Peritrich)
var _ IDer = new(Peritrich)
var _ StartCodoner = new(Peritrich)
var _ StopCodoner = new(Peritrich)

type (
	// Peritrich is the peritrich DNA to protein translation table
	Peritrich struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s Peritrich) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'E', "TAG": 'E',
		"TGT": 'C', "TGC": 'C', "TGG": 'W', "CTT": 'L',
		"CTC": 'L', "CTA": 'L', "CTG": 'L', "CCT": 'P',
		"CCC": 'P', "CCA": 'P', "CCG": 'P', "CAT": 'H',
		"CAC": 'H', "CAA": 'Q', "CAG": 'Q', "CGT": 'R',
		"CGC": 'R', "CGA": 'R', "CGG": 'R', "ATT": 'I',
		"ATC": 'I', "ATA": 'I', "ATG": 'M', "ACT": 'T',
		"ACC": 'T', "ACA": 'T', "ACG": 'T', "AAT": 'N',
		"AAC": 'N', "AAA": 'K', "AAG": 'K', "AGT": 'S',
		"AGC": 'S', "AGA": 'R', "AGG": 'R', "GTT": 'V',
		"GTC": 'V', "GTA": 'V', "GTG": 'V', "GCT": 'A',
		"GCC": 'A', "GCA": 'A', "GCG": 'A', "GAT": 'D',
		"GAC": 'D', "GAA": 'E', "GAG": 'E', "GGT": 'G',
		"GGC": 'G', "GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s Peritrich) String() string {
	return "Peritrich Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s Peritrich) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s Peritrich) ID() uint {
	return 30
}

// StartCodons lists the codons which start a transcript
func (s Peritrich) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s Peritrich) StopCodons() []string {
	return []string{"TGA"}
}
