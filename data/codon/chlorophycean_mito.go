package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(ChlorophyceanMt)
var _ Translater = new(ChlorophyceanMt)
var _ AltNamer = new(ChlorophyceanMt)
var _ IDer = new(ChlorophyceanMt)
var _ StartCodoner = new(ChlorophyceanMt)
var _ StopCodoner = new(ChlorophyceanMt)

type (
	// ChlorophyceanMt is the chlorophycean mtDNA to protein translation table
	ChlorophyceanMt struct{}
)

// Translate converts a codon into its amino acid equivalen
func (s ChlorophyceanMt) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAG": 'L', "TGT": 'C',
		"TGC": 'C', "TGG": 'W', "CTT": 'L', "CTC": 'L',
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
func (s ChlorophyceanMt) String() string {
	return "Chlorophycean Mitochondrial Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s ChlorophyceanMt) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s ChlorophyceanMt) ID() uint {
	return 16
}

// StartCodons lists the codons which start a transcrip
func (s ChlorophyceanMt) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcrip
func (s ChlorophyceanMt) StopCodons() []string {
	return []string{"TAA", "TGA"}
}
