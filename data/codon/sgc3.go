package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC3)
var _ Translater = new(SGC3)
var _ AltNamer = new(SGC3)
var _ StartCodoner = new(SGC3)
var _ StopCodoner = new(SGC3)

type SGC3 struct{}
type MoldMitochondrial SGC3
type ProtozoanMitochondrial SGC3
type CoelenterateMitochondrial SGC3
type Mycoplasma SGC3
type Spiroplasma SGC3

func (s SGC3) Translate(c string) (byte, bool) {
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
		"AGA": 'R', "AGG": 'R', "GTT": 'V', "GTC": 'V',
		"GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A',
		"GCA": 'A', "GCG": 'A', "GAT": 'D', "GAC": 'D',
		"GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G',
		"GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

func (s SGC3) String() string {
	return "SGC3 Codon Library"
}

func (s MoldMitochondrial) String() string {
	return "Mold Mitochondrial Codon Library"
}

func (s ProtozoanMitochondrial) String() string {
	return "Protozoan Mitochondrial Codon Library"
}

func (s CoelenterateMitochondrial) String() string {
	return "Coelenterate Mitochondrial Codon Library"
}

func (s Mycoplasma) String() string {
	return "Mycoplasma Codon Library"
}

func (s Spiroplasma) String() string {
	return "Spiroplasmal Codon Library"
}

func (s SGC3) AltName() string {
	return "SGC3"
}

func (s SGC3) StartCodons() []string {
	return []string{"TTA", "TTG", "CTG", "ATT",
		"ATC", "ATA", "ATG", "GTG"}
}

func (s SGC3) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
