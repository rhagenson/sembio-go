package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(Peritrich)
var _ Translater = new(Peritrich)
var _ AltNamer = new(Peritrich)
var _ StartCodoner = new(Peritrich)
var _ StopCodoner = new(Peritrich)

type Peritrich struct{}

func (s Peritrich) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'E', "TAG": 'E',
		"TGT": 'C', "TGC": 'C', "TGG": 'W', "CTT": 'L',
		"CTC": 'L', "CTA": 'L', "CTG": 'L', "CCT": 'P',
		"CCC": 'P', "CCA": 'P', "CCG": 'P', "CAT": 'H',
		"CAC": 'H', "CAA": 'Q', "CAG": 'Q', "CGT": 'R',
		"CGC": 'R', "CGA": 'R', "CGG": 'R', "ATT": 'I',
		"ATC": 'I', "ATA": 'I', "ATG": 'M', "ACT": 'T',
		"ACC": 'T', "ACA": 'T', "ACG": 'T', "AAT": 'N',
		"AAC": 'N', "AAA": 'K', "AAG": 'K', "AGT": 'S',
		"AGC": 'S', "AGA": 'R', "AGG": 'R', "GTT": 'V',
		"GTC": 'V', "GTA": 'V', "GTG": 'V', "GCT": 'A',
		"GCC": 'A', "GCA": 'A', "GCG": 'A', "GAT": 'D',
		"GAC": 'D', "GAA": 'E', "GAG": 'E', "GGT": 'G',
		"GGC": 'G', "GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

func (s Peritrich) String() string {
	return "Peritrich Nuclear Codon Library"
}

func (s Peritrich) AltName() string {
	return ""
}

func (s Peritrich) StartCodons() []string {
	return []string{"ATG"}
}

func (s Peritrich) StopCodons() []string {
	return []string{"TGA"}
}
