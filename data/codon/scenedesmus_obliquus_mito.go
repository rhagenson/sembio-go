package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(ScenedesmusObliquusMitochondrial)
var _ Translater = new(ScenedesmusObliquusMitochondrial)
var _ AltNamer = new(ScenedesmusObliquusMitochondrial)
var _ StartCodoner = new(ScenedesmusObliquusMitochondrial)
var _ StopCodoner = new(ScenedesmusObliquusMitochondrial)

type ScenedesmusObliquusMitochondrial struct{}

func (s ScenedesmusObliquusMitochondrial) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCG": 'S', "TAT": 'Y',
		"TAC": 'Y', "TAG": 'L', "TGT": 'C', "TGC": 'C',
		"TGG": 'W', "CTT": 'L', "CTC": 'L', "CTA": 'L',
		"CTG": 'L', "CCT": 'P', "CCC": 'P', "CCA": 'P',
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

func (s ScenedesmusObliquusMitochondrial) String() string {
	return "Scenedesmus obliquus Mitochondrial Codon Library"
}

func (s ScenedesmusObliquusMitochondrial) AltName() string {
	return ""
}

func (s ScenedesmusObliquusMitochondrial) StartCodons() []string {
	return []string{"ATG"}
}

func (s ScenedesmusObliquusMitochondrial) StopCodons() []string {
	return []string{"TCA", "TAA", "TGA"}
}
