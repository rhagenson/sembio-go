package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC3)
var _ Translater = new(SGC3)
var _ AltNamer = new(SGC3)
var _ IDer = new(SGC3)
var _ StartCodoner = new(SGC3)
var _ StopCodoner = new(SGC3)

type (
	// SGC3 is the NCBI mtDNA to protein translation table for
	// Mold, Protozoan, and Coelenterate
	// As well as the standard code for Mycoplasma/Spiroplasma
	SGC3 struct{}

	// MoldMt is the mold mtDNA to protein translation table
	MoldMt SGC3

	// ProtozoanMt is the protozoan mtDNA to protein translation table
	ProtozoanMt SGC3

	// CoelenterateMt is the coelenterate mtDNA to protein translation table
	CoelenterateMt SGC3

	// Mycoplasma is the mycoplasma DNA to protein translation table
	Mycoplasma SGC3

	// Spiroplasma is the spiroplasma DNA to protein translation table
	Spiroplasma SGC3
)

// Translate converts a codon into its amino acid equivalent
func (s SGC3) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGA": 'W', "TGG": 'W', "CTT": 'L', "CTC": 'L',
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
func (s SGC3) String() string {
	return "SGC3 Codon Library"
}

// String provides a human-readable indication of usage
func (s MoldMt) String() string {
	return "Mold Mitochondrial Codon Library"
}

// String provides a human-readable indication of usage
func (s ProtozoanMt) String() string {
	return "Protozoan Mitochondrial Codon Library"
}

// String provides a human-readable indication of usage
func (s CoelenterateMt) String() string {
	return "Coelenterate Mitochondrial Codon Library"
}

// String provides a human-readable indication of usage
func (s Mycoplasma) String() string {
	return "Mycoplasma Codon Library"
}

// String provides a human-readable indication of usage
func (s Spiroplasma) String() string {
	return "Spiroplasma Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s SGC3) AltName() string {
	return "SGC3"
}

// ID provides the alternative identifier used by NCBI
func (s SGC3) ID() uint {
	return 4
}

// StartCodons lists the codons which start a transcript
func (s SGC3) StartCodons() []string {
	return []string{"TTA", "TTG", "CTG", "ATT", "ATC", "ATA", "ATG", "GTG"}
}

// StopCodons lists the codons which end a transcript
func (s SGC3) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
