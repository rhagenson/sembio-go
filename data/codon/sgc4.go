package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC4)
var _ Translater = new(SGC4)
var _ AltNamer = new(SGC4)
var _ StartCodoner = new(SGC4)
var _ StopCodoner = new(SGC4)

type SGC4 struct{}
type InvertebrateMitochondrial SGC4

func (s SGC4) Translate(c string) (byte, bool) {
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
		"AGA": 'S', "AGG": 'S', "GTT": 'V', "GTC": 'V',
		"GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A',
		"GCA": 'A', "GCG": 'A', "GAT": 'D', "GAC": 'D',
		"GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G',
		"GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

func (s SGC4) String() string {
	return "SGC4 Codon Library"
}

func (s InvertebrateMitochondrial) String() string {
	return "Invertebrate Mitrochondrial Codon Library"
}

func (s SGC4) AltName() string {
	return "SGC4"
}

func (s SGC4) StartCodons() []string {
	return []string{"TTG", "ATT", "ATC", "ATA", "ATG", "GTG"}
}

func (s SGC4) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
