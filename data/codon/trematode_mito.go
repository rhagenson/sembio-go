package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(TrematodeMt)
var _ Translater = new(TrematodeMt)
var _ AltNamer = new(TrematodeMt)
var _ IDer = new(TrematodeMt)
var _ StartCodoner = new(TrematodeMt)
var _ StopCodoner = new(TrematodeMt)

type (
	// TrematodeMt is the trematode mtDNA to protein translation table
	TrematodeMt struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s TrematodeMt) Translate(c string) (byte, bool) {
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
		"AAA": 'N', "AAG": 'K', "AGT": 'S', "AGC": 'S',
		"AGA": 'S', "AGG": 'S', "GTT": 'V', "GTC": 'V',
		"GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A',
		"GCA": 'A', "GCG": 'A', "GAT": 'D', "GAC": 'D',
		"GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G',
		"GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s TrematodeMt) String() string {
	return "Trematode Mitochondrial Codon Library"
}

// AltName provides the alternative name used by NCB
func (s TrematodeMt) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s TrematodeMt) ID() uint {
	return 21
}

// StartCodons lists the codons which start a transcript
func (s TrematodeMt) StartCodons() []string {
	return []string{"ATG", "GTG"}
}

// StopCodons lists the codons which end a transcript
func (s TrematodeMt) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
