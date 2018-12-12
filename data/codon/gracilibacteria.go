package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(Gracilibacteria)
var _ Translater = new(Gracilibacteria)
var _ AltNamer = new(Gracilibacteria)
var _ StartCodoner = new(Gracilibacteria)
var _ StopCodoner = new(Gracilibacteria)

type Gracilibacteria struct{}
type CandidateDivisionSR1 Gracilibacteria

func (s Gracilibacteria) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
		"TGA": 'G', "TGG": 'W', "CTT": 'L', "CTC": 'L',
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

func (s Gracilibacteria) String() string {
	return "Gracilibacteria Codon Library"
}

func (s CandidateDivisionSR1) String() string {
	return "Candidate Division SR1 Codon Library"
}

func (s Gracilibacteria) AltName() string {
	return ""
}

func (s Gracilibacteria) StartCodons() []string {
	return []string{"TTG", "ATG", "GTG"}
}

func (s Gracilibacteria) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
