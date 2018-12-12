package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC9)
var _ Translater = new(SGC9)
var _ AltNamer = new(SGC9)
var _ IDer = new(SGC9)
var _ StartCodoner = new(SGC9)
var _ StopCodoner = new(SGC9)

type (
	// SGC9 is the NCBI mtDNA to protein translation table for
	// Echinoderm and Flatworm
	SGC9 struct{}

	// Euplotid is the euplotid DNA to protein translation table
	Euplotid SGC9
)

// Translate converts a codon into its amino acid equivalent
func (s SGC9) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGA": 'C', "TGG": 'W', "CTT": 'L', "CTC": 'L',
		"CTA": 'L', "CTG": 'L', "CCT": 'P', "CCC": 'P',
		"CCA": 'P', "CCG": 'P', "CAT": 'H', "CAC": 'H',
		"CAA": 'Q', "CAG": 'Q', "CGT": 'R', "CGC": 'R',
		"CGA": 'R', "CGG": 'R', "ATT": 'I', "ATC": 'I',
		"ATA": 'I', "ATG": 'M', "ACT": 'T', "ACC": 'T',
		"ACA": 'T', "ACG": 'T', "AAT": 'N', "AAC": 'N',
		"AAA": 'K', "AAG": 'K', "AGT": 'S', "AGC": 'S',
		"AGA": 'R', "AGG": 'R', "GTT": 'V', "GTC": 'V',
		"GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A',
		"GCA": 'A', "GCG": 'A', "GAT": 'D', "GAC": 'D',
		"GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G',
		"GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s SGC9) String() string {
	return "SGC9 Codon Library"
}

// String provides a human-readable indication of usage
func (s Euplotid) String() string {
	return "Euplotid Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s SGC9) AltName() string {
	return "SGC9"
}

// ID provides the alternative identifier used by NCBI
func (s SGC9) ID() uint {
	return 10
}

// StartCodons lists the codons which start a transcript
func (s SGC9) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s SGC9) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
