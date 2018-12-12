package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(ThraustochytriumMitochondrial)
var _ Translater = new(ThraustochytriumMitochondrial)
var _ AltNamer = new(ThraustochytriumMitochondrial)
var _ StartCodoner = new(ThraustochytriumMitochondrial)
var _ StopCodoner = new(ThraustochytriumMitochondrial)

type ThraustochytriumMitochondrial struct{}

func (s ThraustochytriumMitochondrial) Translate(c string) (byte, bool) {
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

func (s ThraustochytriumMitochondrial) String() string {
	return "Thraustochytrium Mitochondrial Codon Library"
}

func (s ThraustochytriumMitochondrial) AltName() string {
	return ""
}

func (s ThraustochytriumMitochondrial) StartCodons() []string {
	return []string{"ATT", "ATG", "GTG"}
}

func (s ThraustochytriumMitochondrial) StopCodons() []string {
	return []string{"TTA", "TAA", "TAG", "TGA"}
}
