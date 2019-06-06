package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(ThraustochytriumMt)
var _ Translater = new(ThraustochytriumMt)
var _ AltNamer = new(ThraustochytriumMt)
var _ IDer = new(ThraustochytriumMt)
var _ StartCodoner = new(ThraustochytriumMt)
var _ StopCodoner = new(ThraustochytriumMt)

type (
	// ThraustochytriumMt is the thraustochytrium mtDNA to protein translation table
	ThraustochytriumMt struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s ThraustochytriumMt) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTG": 'L', "TCT": 'S',
		"TCC": 'S', "TCA": 'S', "TCG": 'S', "TAT": 'Y',
		"TAC": 'Y', "TGT": 'C', "TGC": 'C', "TGG": 'W',
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
func (s ThraustochytriumMt) String() string {
	return "Thraustochytrium Mitochondrial Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s ThraustochytriumMt) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s ThraustochytriumMt) ID() uint {
	return 23
}

// StartCodons lists the codons which start a transcript
func (s ThraustochytriumMt) StartCodons() []string {
	return []string{"ATT", "ATG", "GTG"}
}

// StopCodons lists the codons which end a transcript
func (s ThraustochytriumMt) StopCodons() []string {
	return []string{"TTA", "TAA", "TAG", "TGA"}
}
