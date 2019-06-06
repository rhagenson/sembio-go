package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(Mesodinium)
var _ Translater = new(Mesodinium)
var _ AltNamer = new(Mesodinium)
var _ IDer = new(Mesodinium)
var _ StartCodoner = new(Mesodinium)
var _ StopCodoner = new(Mesodinium)

type (
	// Mesodinium is the mesodinium DNA to protein translation table
	Mesodinium struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s Mesodinium) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'Y', "TAG": 'Y',
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
func (s Mesodinium) String() string {
	return "Mesodinium Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s Mesodinium) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s Mesodinium) ID() uint {
	return 29
}

// StartCodons lists the codons which start a transcript
func (s Mesodinium) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s Mesodinium) StopCodons() []string {
	return []string{"TGA"}
}
