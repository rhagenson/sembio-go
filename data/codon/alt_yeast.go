package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(AlternativeYeastNuclear)
var _ Translater = new(AlternativeYeastNuclear)
var _ AltNamer = new(AlternativeYeastNuclear)
var _ StartCodoner = new(AlternativeYeastNuclear)
var _ StopCodoner = new(AlternativeYeastNuclear)

type AlternativeYeastNuclear struct{}

func (s AlternativeYeastNuclear) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGG": 'W', "CTT": 'L', "CTC": 'L', "CTA": 'L',
		"CTG": 'S', "CCT": 'P', "CCC": 'P', "CCA": 'P',
		"CCG": 'P', "CAT": 'H', "CAC": 'H', "CAA": 'Q',
		"CAG": 'Q', "CGT": 'R', "CGC": 'R', "CGA": 'R',
		"CGG": 'R', "ATT": 'I', "ATC": 'I', "ATA": 'I',
		"ATG": 'M', "ACT": 'T', "ACC": 'T', "ACA": 'T',
		"ACG": 'T', "AAT": 'N', "AAC": 'N', "AAA": 'K',
		"AAG": 'K', "AGT": 'S', "AGC": 'S', "AGA": 'R',
		"AGG": 'R', "GTT": 'V', "GTC": 'V', "GTA": 'V',
		"GTG": 'V', "GCT": 'A', "GCC": 'A', "GCA": 'A',
		"GCG": 'A', "GAT": 'D', "GAC": 'D', "GAA": 'E',
		"GAG": 'E', "GGT": 'G', "GGC": 'G', "GGA": 'G',
		"GGG": 'G',
	}[c]
	return aa, ok
}

func (s AlternativeYeastNuclear) String() string {
	return "Alternative Yeast Nuclear Codon Library"
}

func (s AlternativeYeastNuclear) AltName() string {
	return ""
}

func (s AlternativeYeastNuclear) StartCodons() []string {
	return []string{"CTG", "ATG"}
}

func (s AlternativeYeastNuclear) StopCodons() []string {
	return []string{"TAA", "TAG", "TGA"}
}
