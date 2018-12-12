package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC5)
var _ Translater = new(SGC5)
var _ AltNamer = new(SGC5)
var _ IDer = new(SGC5)
var _ StartCodoner = new(SGC5)
var _ StopCodoner = new(SGC5)

type (
	// SGC5 is the NCBI DNA to protein translation table for
	// Ciliate, Dasycladacean, and Hexamita
	SGC5 struct{}

	// Ciliate is the DNA to protein translation table
	Ciliate SGC5

	// Dasycladacean is the DNA to protein translation table
	Dasycladacean SGC5

	// Hexamita is the DNA to protein translation table
	Hexamita SGC5
)

// Translate converts a codon into its amino acid equivalent
func (s SGC5) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'Q', "TAG": 'Q',
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

// String provides a human-readable indication of usag
func (s SGC5) String() string {
	return "SGC5 Codon Library"
}

// String provides a human-readable indication of usag
func (s Ciliate) String() string {
	return "Ciliate Codon Library"
}

// String provides a human-readable indication of usag
func (s Dasycladacean) String() string {
	return "Dasycladacean Codon Library"
}

// String provides a human-readable indication of usag
func (s Hexamita) String() string {
	return "Hexamita Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s SGC5) AltName() string {
	return "SGC5"
}

// ID provides the alternative identifier used by NCBI
func (s SGC5) ID() uint {
	return 6
}

// StartCodons lists the codons which start a transcript
func (s SGC5) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s SGC5) StopCodons() []string {
	return []string{"TGA"}
}
