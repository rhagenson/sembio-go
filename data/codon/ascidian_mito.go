package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(AscidianMitochondrial)
var _ Translater = new(AscidianMitochondrial)
var _ AltNamer = new(AscidianMitochondrial)
var _ StartCodoner = new(AscidianMitochondrial)
var _ StopCodoner = new(AscidianMitochondrial)

type AscidianMitochondrial struct{}

func (s AscidianMitochondrial) Translate(c string) (byte, bool) {
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

func (s AscidianMitochondrial) String() string {
	return "Ascidian Mitochondrial Codon Library"
}

func (s AscidianMitochondrial) AltName() string {
	return ""
}

func (s AscidianMitochondrial) StartCodons() []string {
	return []string{"TTG", "ATA", "ATG", "GTG"}
}

func (s AscidianMitochondrial) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
