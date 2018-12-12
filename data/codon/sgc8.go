package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(SGC8)
var _ Translater = new(SGC8)
var _ AltNamer = new(SGC8)
var _ StartCodoner = new(SGC8)
var _ StopCodoner = new(SGC8)

type SGC8 struct{}
type EchinodermMitochondrial SGC8
type FlatwormMitochondrial SGC8

func (s SGC8) Translate(c string) (byte, bool) {
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

func (s SGC8) String() string {
	return "SGC8 Codon Library"
}

func (s EchinodermMitochondrial) String() string {
	return "Echinoderm Mitochondrial Codon Library"
}

func (s FlatwormMitochondrial) String() string {
	return "Flatworm Mitochondrial Codon Library"
}

func (s SGC8) AltName() string {
	return "SGC8"
}

func (s SGC8) StartCodons() []string {
	return []string{"ATG", "GTG"}
}

func (s SGC8) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
