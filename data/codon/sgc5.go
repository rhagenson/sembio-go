package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC5)
var _ Translater = new(SGC5)
var _ AltNamer = new(SGC5)
var _ StartCodoner = new(SGC5)
var _ StopCodoner = new(SGC5)

type SGC5 struct{}
type CiliateNuclear SGC5
type DasycladaceanNuclear SGC5
type HexamitaNuclear SGC5

func (s SGC5) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'Q', "TAG": 'Q',
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

func (s SGC5) String() string {
	return "SGC5 Codon Library"
}

func (s CiliateNuclear) String() string {
	return "Ciliate Nuclear Codon Library"
}

func (s DasycladaceanNuclear) String() string {
	return "Dasycladacean Nuclear Codon Library"
}

func (s HexamitaNuclear) String() string {
	return "Hexamita Nuclear Codon Library"
}

func (s SGC5) AltName() string {
	return "SGC5"
}

func (s SGC5) StartCodons() []string {
	return []string{"ATG"}
}

func (s SGC5) StopCodons() []string {
	return []string{"TGA"}
}
