package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC2)
var _ Translater = new(SGC2)
var _ AltNamer = new(SGC2)
var _ StartCodoner = new(SGC2)
var _ StopCodoner = new(SGC2)

type SGC2 struct{}
type YeastMitochondrial SGC2

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

func (s SGC2) String() string {
	return "SGC2 Codon Library"
}

func (s YeastMitochondrial) String() string {
	return "Yeast Mitrochondrial Codon Library"
}

func (s SGC2) AltName() string {
	return "SGC2"
}

func (s SGC2) StartCodons() []string {
	return []string{"ATA", "ATG"}
}

func (s SGC2) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
