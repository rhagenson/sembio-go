package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(AltFlatwormMt)
var _ Translater = new(AltFlatwormMt)
var _ AltNamer = new(AltFlatwormMt)
var _ IDer = new(AltFlatwormMt)
var _ StartCodoner = new(AltFlatwormMt)
var _ StopCodoner = new(AltFlatwormMt)

type (
	// AltFlatwormMt is the NCBI mtDNA to protein translation table for flatworm
	AltFlatwormMt struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s AltFlatwormMt) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'Y', "TGT": 'C',
		"TGC": 'C', "TGA": 'W', "TGG": 'W', "CTT": 'L',
		"CTC": 'L', "CTA": 'L', "CTG": 'L', "CCT": 'P',
		"CCC": 'P', "CCA": 'P', "CCG": 'P', "CAT": 'H',
		"CAC": 'H', "CAA": 'Q', "CAG": 'Q', "CGT": 'R',
		"CGC": 'R', "CGA": 'R', "CGG": 'R', "ATT": 'I',
		"ATC": 'I', "ATA": 'I', "ATG": 'M', "ACT": 'T',
		"ACC": 'T', "ACA": 'T', "ACG": 'T', "AAT": 'N',
		"AAC": 'N', "AAA": 'N', "AAG": 'K', "AGT": 'S',
		"AGC": 'S', "AGA": 'S', "AGG": 'S', "GTT": 'V',
		"GTC": 'V', "GTA": 'V', "GTG": 'V', "GCT": 'A',
		"GCC": 'A', "GCA": 'A', "GCG": 'A', "GAT": 'D',
		"GAC": 'D', "GAA": 'E', "GAG": 'E', "GGT": 'G',
		"GGC": 'G', "GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s AltFlatwormMt) String() string {
	return "Alternative Flatworm Mt Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s AltFlatwormMt) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s AltFlatwormMt) ID() uint {
	return 14
}

// StartCodons lists the codons which start a transcript
func (s AltFlatwormMt) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s AltFlatwormMt) StopCodons() []string {
	return []string{"TAG"}
}
