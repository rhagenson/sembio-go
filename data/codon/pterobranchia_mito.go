package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(PterobranchiaMitochondrial)
var _ Translater = new(PterobranchiaMitochondrial)
var _ AltNamer = new(PterobranchiaMitochondrial)
var _ StartCodoner = new(PterobranchiaMitochondrial)
var _ StopCodoner = new(PterobranchiaMitochondrial)

type PterobranchiaMitochondrial struct{}

func (s PterobranchiaMitochondrial) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGA": 'W', "TGG": 'W', "CTT": 'L', "CTC": 'L',
		"CTA": 'L', "CTG": 'L', "CCT": 'P', "CCC": 'P',
		"CCA": 'P', "CCG": 'P', "CAT": 'H', "CAC": 'H',
		"CAA": 'Q', "CAG": 'Q', "CGT": 'R', "CGC": 'R',
		"CGA": 'R', "CGG": 'R', "ATT": 'I', "ATC": 'I',
		"ATA": 'I', "ATG": 'M', "ACT": 'T', "ACC": 'T',
		"ACA": 'T', "ACG": 'T', "AAT": 'N', "AAC": 'N',
		"AAA": 'K', "AAG": 'K', "AGT": 'S', "AGC": 'S',
		"AGA": 'S', "AGG": 'K', "GTT": 'V', "GTC": 'V',
		"GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A',
		"GCA": 'A', "GCG": 'A', "GAT": 'D', "GAC": 'D',
		"GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G',
		"GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

func (s PterobranchiaMitochondrial) String() string {
	return "Pterobranchia Mitochondrial Codon Library"
}

func (s PterobranchiaMitochondrial) AltName() string {
	return ""
}

func (s PterobranchiaMitochondrial) StartCodons() []string {
	return []string{"TTG", "CTG", "ATG", "GTG"}
}

func (s PterobranchiaMitochondrial) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
