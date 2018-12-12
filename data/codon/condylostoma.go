package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(Condylostoma)
var _ Translater = new(Condylostoma)
var _ AltNamer = new(Condylostoma)
var _ IDer = new(Condylostoma)
var _ StartCodoner = new(Condylostoma)
var _ StopCodoner = new(Condylostoma)

type (
	// Condylostoma is the condylostoma DNA to protein translation table
	Condylostoma struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s Condylostoma) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'Q', "TAG": 'Q',
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
func (s Condylostoma) String() string {
	return "Condylostoma Codon Library"
}

// AltName provides the alternative name used by NCB
func (s Condylostoma) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s Condylostoma) ID() uint {
	return 28
}

// StartCodons lists the codons which start a transcript
func (s Condylostoma) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s Condylostoma) StopCodons() []string {
	return []string{"TAA", "TAG", "TGA"}
}
