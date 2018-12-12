package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(ScenedesmusObliquusMt)
var _ Translater = new(ScenedesmusObliquusMt)
var _ AltNamer = new(ScenedesmusObliquusMt)
var _ IDer = new(ScenedesmusObliquusMt)
var _ StartCodoner = new(ScenedesmusObliquusMt)
var _ StopCodoner = new(ScenedesmusObliquusMt)

type (
	// ScenedesmusObliquusMt is the Scenedesmus obliquus mtDNA to protein translation tabl
	ScenedesmusObliquusMt struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s ScenedesmusObliquusMt) Translate(c string) (byte, bool) {
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

// String provides a human-readable indication of usage
func (s ScenedesmusObliquusMt) String() string {
	return "Scenedesmus obliquus Mitochondrial Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s ScenedesmusObliquusMt) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s ScenedesmusObliquusMt) ID() uint {
	return 22
}

// StartCodons lists the codons which start a transcript
func (s ScenedesmusObliquusMt) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s ScenedesmusObliquusMt) StopCodons() []string {
	return []string{"TCA", "TAA", "TGA"}
}
