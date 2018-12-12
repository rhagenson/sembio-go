package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC1)
var _ Translater = new(SGC1)
var _ AltNamer = new(SGC1)
var _ StartCodoner = new(SGC1)
var _ StopCodoner = new(SGC1)

type SGC1 struct{}
type VertebrateMitochondrial SGC1

func (s SGC1) Translate(c string) (byte, bool) {
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
		"GTT": 'V', "GTC": 'V', "GTA": 'V', "GTG": 'V',
		"GCT": 'A', "GCC": 'A', "GCA": 'A', "GCG": 'A',
		"GAT": 'D', "GAC": 'D', "GAA": 'E', "GAG": 'E',
		"GGT": 'G', "GGC": 'G', "GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

func (s SGC1) String() string {
	return "SGC1 Codon Library"
}

func (s VertebrateMitochondrial) String() string {
	return "Vertebrate Mitochondrial Codon Library"
}

func (s SGC1) AltName() string {
	return "SGC1"
}

func (s SGC1) StartCodons() []string {
	return []string{"ATT", "ATC", "ATA", "ATG", "GTG"}
}

func (s SGC1) StopCodons() []string {
	return []string{"TAA", "TAG", "AGA", "AGG"}
}
