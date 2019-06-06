package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC1)
var _ Translater = new(SGC1)
var _ AltNamer = new(SGC1)
var _ IDer = new(SGC1)
var _ StartCodoner = new(SGC1)
var _ StopCodoner = new(SGC1)

type (
	// SGC1 is the NCBI mtDNA to protein translation table
	SGC1 struct{}

	// VertebrateMt is the mtDNA to protein translation table
	VertebrateMt struct {
		SGC1
	}
)

// Translate converts a codon into its amino acid equivalent
func (s SGC1) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGA": 'W', "TGG": 'W', "CTT": 'L', "CTC": 'L',
		"CTA": 'L', "CTG": 'L', "CCT": 'P', "CCC": 'P',
		"CCA": 'P', "CCG": 'P', "CAT": 'H', "CAC": 'H',
		"CAA": 'Q', "CAG": 'Q', "CGT": 'R', "CGC": 'R',
		"CGA": 'R', "CGG": 'R', "ATT": 'I', "ATC": 'I',
		"ATA": 'M', "ATG": 'M', "ACT": 'T', "ACC": 'T',
		"ACA": 'T', "ACG": 'T', "AAT": 'N', "AAC": 'N',
		"AAA": 'K', "AAG": 'K', "AGT": 'S', "AGC": 'S',
		"GTT": 'V', "GTC": 'V', "GTA": 'V', "GTG": 'V',
		"GCT": 'A', "GCC": 'A', "GCA": 'A', "GCG": 'A',
		"GAT": 'D', "GAC": 'D', "GAA": 'E', "GAG": 'E',
		"GGT": 'G', "GGC": 'G', "GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s SGC1) String() string {
	return "SGC1 Codon Library"
}

// String provides a human-readable indication of usage
func (s VertebrateMt) String() string {
	return "Vertebrate Mitochondrial Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s SGC1) AltName() string {
	return "SGC1"
}

// ID provides the alternative identifier used by NCBI
func (s SGC1) ID() uint {
	return 2
}

// StartCodons lists the codons which start a transcript
func (s SGC1) StartCodons() []string {
	return []string{"ATT", "ATC", "ATA", "ATG", "GTG"}
}

// StopCodons lists the codons which end a transcript
func (s SGC1) StopCodons() []string {
	return []string{"TAA", "TAG", "AGA", "AGG"}
}
