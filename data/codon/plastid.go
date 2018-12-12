package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(BacterialArchaelPlantPlastid)
var _ Translater = new(BacterialArchaelPlantPlastid)
var _ AltNamer = new(BacterialArchaelPlantPlastid)
var _ StartCodoner = new(BacterialArchaelPlantPlastid)
var _ StopCodoner = new(BacterialArchaelPlantPlastid)

type BacterialArchaelPlantPlastid struct{}
type BacterialPlastid BacterialArchaelPlantPlastid
type ArchaelPlastid BacterialArchaelPlantPlastid
type PlantPlastid BacterialArchaelPlantPlastid

func (s BacterialArchaelPlantPlastid) Translate(c string) (byte, bool) {
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

func (s BacterialArchaelPlantPlastid) String() string {
	return "Bacterial, Archael, and Plant Plastid Codon Library"
}

func (s BacterialPlastid) String() string {
	return "Bacterial Plastid Codon Library"
}

func (s ArchaelPlastid) String() string {
	return "Archael Plastid Codon Library"
}

func (s PlantPlastid) String() string {
	return "Plant Plastid Codon Library"
}

func (s BacterialArchaelPlantPlastid) AltName() string {
	return ""
}

func (s BacterialArchaelPlantPlastid) StartCodons() []string {
	return []string{"TTG", "CTG", "ATT",
		"ATC", "ATA", "ATG", "GTG"}
}

func (s BacterialArchaelPlantPlastid) StopCodons() []string {
	return []string{"TAA", "TAG", "TGA"}
}
