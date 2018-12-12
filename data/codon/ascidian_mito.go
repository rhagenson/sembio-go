package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(AscidianMt)
var _ Translater = new(AscidianMt)
var _ AltNamer = new(AscidianMt)
var _ IDer = new(AscidianMt)
var _ StartCodoner = new(AscidianMt)
var _ StopCodoner = new(AscidianMt)

type (
	// AscidianMt is the ascidian mtDNA to protein translation table
	AscidianMt struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s AscidianMt) Translate(c string) (byte, bool) {
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
		"AGA": 'G', "AGG": 'G', "GTT": 'V', "GTC": 'V',
		"GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A',
		"GCA": 'A', "GCG": 'A', "GAT": 'D', "GAC": 'D',
		"GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G',
		"GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usage
func (s AscidianMt) String() string {
	return "Ascidian Mitochondrial Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s AscidianMt) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s AscidianMt) ID() uint {
	return 13
}

// StartCodons lists the codons which start a transcript
func (s AscidianMt) StartCodons() []string {
	return []string{"TTG", "ATA", "ATG", "GTG"}
}

// StopCodons lists the codons which end a transcript
func (s AscidianMt) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
