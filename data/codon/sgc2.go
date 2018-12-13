package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC2)
var _ Translater = new(SGC2)
var _ AltNamer = new(SGC2)
var _ IDer = new(SGC2)
var _ StartCodoner = new(SGC2)
var _ StopCodoner = new(SGC2)

type (
	// SGC2 is the NCBI yeast mtDNA to protein translation table
	SGC2 struct{}

	// YeastMt is the yeast mtDNA to protein translation table
	YeastMt struct {
		SGC2
	}
)

// Translate converts a codon into its amino acid equivalent
func (s SGC2) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGA": 'W', "TGG": 'W', "CTT": 'T', "CTC": 'T',
		"CTA": 'T', "CTG": 'T', "CCT": 'P', "CCC": 'P',
		"CCA": 'P', "CCG": 'P', "CAT": 'H', "CAC": 'H',
		"CAA": 'Q', "CAG": 'Q', "CGT": 'R', "CGC": 'R',
		"CGA": 'R', "CGG": 'R', "ATT": 'I', "ATC": 'I',
		"ATA": 'M', "ATG": 'M', "ACT": 'T', "ACC": 'T',
		"ACA": 'T', "ACG": 'T', "AAT": 'N', "AAC": 'N',
		"AAA": 'K', "AAG": 'K', "AGT": 'S', "AGC": 'S',
		"AGA": 'R', "AGG": 'R', "GTT": 'V', "GTC": 'V',
		"GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A',
		"GCA": 'A', "GCG": 'A', "GAT": 'D', "GAC": 'D',
		"GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G',
		"GGA": 'G', "GGG": 'G', "TCT": 'S', "TCC": 'S',
		"TCA": 'S', "TCG": 'S',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s SGC2) String() string {
	return "SGC2 Codon Library"
}

// String provides a human-readable indication of usage
func (s YeastMt) String() string {
	return "Yeast Mitochondrial Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s SGC2) AltName() string {
	return "SGC2"
}

// ID provides the alternative identifier used by NCBI
func (s SGC2) ID() uint {
	return 3
}

// StartCodons lists the codons which start a transcript
func (s SGC2) StartCodons() []string {
	return []string{"ATA", "ATG"}
}

// StopCodons lists the codons which end a transcript
func (s SGC2) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
