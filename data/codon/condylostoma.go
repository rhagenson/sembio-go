package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(Condylostoma)
var _ Translater = new(Condylostoma)
var _ AltNamer = new(Condylostoma)
var _ StartCodoner = new(Condylostoma)
var _ StopCodoner = new(Condylostoma)

type Condylostoma struct{}

func (s Condylostoma) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'Q', "TAG": 'Q',
		"TGT": 'C', "TGC": 'C', "TGA": 'W', "TGG": 'W',
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

func (s Condylostoma) String() string {
	return "Condylostoma Nuclear Codon Library"
}

func (s Condylostoma) AltName() string {
	return ""
}

func (s Condylostoma) StartCodons() []string {
	return []string{"ATG"}
}

func (s Condylostoma) StopCodons() []string {
	return []string{"TAA", "TAG", "TGA"}
}
