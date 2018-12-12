package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(TrematodeMitochondrial)
var _ Translater = new(TrematodeMitochondrial)
var _ AltNamer = new(TrematodeMitochondrial)
var _ StartCodoner = new(TrematodeMitochondrial)
var _ StopCodoner = new(TrematodeMitochondrial)

type TrematodeMitochondrial struct{}

func (s TrematodeMitochondrial) Translate(c string) (byte, bool) {
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

func (s TrematodeMitochondrial) String() string {
	return "Trematode Mitochondrial Codon Library"
}

func (s TrematodeMitochondrial) AltName() string {
	return ""
}

func (s TrematodeMitochondrial) StartCodons() []string {
	return []string{"ATG", "GTG"}
}

func (s TrematodeMitochondrial) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
