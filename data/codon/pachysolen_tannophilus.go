package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(PachysolenTannophilus)
var _ Translater = new(PachysolenTannophilus)
var _ AltNamer = new(PachysolenTannophilus)
var _ StartCodoner = new(PachysolenTannophilus)
var _ StopCodoner = new(PachysolenTannophilus)

type PachysolenTannophilus struct{}

func (s PachysolenTannophilus) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGG": 'W', "CTT": 'L', "CTC": 'L', "CTA": 'L',
		"CTG": 'A', "CCT": 'P', "CCC": 'P', "CCA": 'P',
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

func (s PachysolenTannophilus) String() string {
	return "Pachysolen tannophilus Nuclear Codon Library"
}
func (s PachysolenTannophilus) AltName() string {
	return ""
}

func (s PachysolenTannophilus) StartCodons() []string {
	return []string{"CTG", "ATG"}
}

func (s PachysolenTannophilus) StopCodons() []string {
	return []string{"TAA", "TAG", "TGA"}
}
