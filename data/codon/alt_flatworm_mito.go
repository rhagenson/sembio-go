package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(AltFlatwormMitrochondrial)
var _ Translater = new(AltFlatwormMitrochondrial)
var _ AltNamer = new(AltFlatwormMitrochondrial)
var _ StartCodoner = new(AltFlatwormMitrochondrial)
var _ StopCodoner = new(AltFlatwormMitrochondrial)

type AltFlatwormMitrochondrial struct{}

func (s AltFlatwormMitrochondrial) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAA": 'Y', "TGT": 'C',
		"TGC": 'C', "TGA": 'W', "TGG": 'W', "CTT": 'L',
		"CTC": 'L', "CTA": 'L', "CTG": 'L', "CCT": 'P',
		"CCC": 'P', "CCA": 'P', "CCG": 'P', "CAT": 'H',
		"CAC": 'H', "CAA": 'Q', "CAG": 'Q', "CGT": 'R',
		"CGC": 'R', "CGA": 'R', "CGG": 'R', "ATT": 'I',
		"ATC": 'I', "ATA": 'I', "ATG": 'M', "ACT": 'T',
		"ACC": 'T', "ACA": 'T', "ACG": 'T', "AAT": 'N',
		"AAC": 'N', "AAA": 'N', "AAG": 'K', "AGT": 'S',
		"AGC": 'S', "AGA": 'S', "AGG": 'S', "GTT": 'V',
		"GTC": 'V', "GTA": 'V', "GTG": 'V', "GCT": 'A',
		"GCC": 'A', "GCA": 'A', "GCG": 'A', "GAT": 'D',
		"GAC": 'D', "GAA": 'E', "GAG": 'E', "GGT": 'G',
		"GGC": 'G', "GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

func (s AltFlatwormMitrochondrial) String() string {
	return "Alternative Flatworm Mitrochondrial Codon Library"
}

func (s AltFlatwormMitrochondrial) AltName() string {
	return ""
}

func (s AltFlatwormMitrochondrial) StartCodons() []string {
	return []string{"ATG"}
}

func (s AltFlatwormMitrochondrial) StopCodons() []string {
	return []string{"TAG"}
}
