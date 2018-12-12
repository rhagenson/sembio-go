package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(Blastocrithidia)
var _ Translater = new(Blastocrithidia)
var _ AltNamer = new(Blastocrithidia)
var _ IDer = new(Blastocrithidia)
var _ StartCodoner = new(Blastocrithidia)
var _ StopCodoner = new(Blastocrithidia)

type (
	// Blastocrithidia is the blastocrithidia DNA to protein translation table
	Blastocrithidia struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s Blastocrithidia) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'E', "TAG": 'E',
		"TGT": 'C', "TGC": 'C', "TGA": 'W', "TGG": 'W',
		"CTT": 'L', "CTC": 'L', "CTA": 'L', "CTG": 'L',
		"CCT": 'P', "CCC": 'P', "CCA": 'P', "CCG": 'P',
		"CAT": 'H', "CAC": 'H', "CAA": 'Q', "CAG": 'Q',
		"CGT": 'R', "CGC": 'R', "CGA": 'R', "CGG": 'R',
		"ATT": 'I', "ATC": 'I', "ATA": 'I', "ATG": 'M',
		"ACT": 'T', "ACC": 'T', "ACA": 'T', "ACG": 'T',
		"AAT": 'N', "AAC": 'N', "AAA": 'K', "AAG": 'K',
		"AGT": 'S', "AGC": 'S', "AGA": 'R', "AGG": 'R',
		"GTT": 'V', "GTC": 'V', "GTA": 'V', "GTG": 'V',
		"GCT": 'A', "GCC": 'A', "GCA": 'A', "GCG": 'A',
		"GAT": 'D', "GAC": 'D', "GAA": 'E', "GAG": 'E',
		"GGT": 'G', "GGC": 'G', "GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s Blastocrithidia) String() string {
	return "Blastocrithidia Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s Blastocrithidia) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s Blastocrithidia) ID() uint {
	return 31
}

// StartCodons lists the codons which start a transcript
func (s Blastocrithidia) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s Blastocrithidia) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
