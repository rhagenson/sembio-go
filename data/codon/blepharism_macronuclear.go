package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(BlepharismaMacronuclear)
var _ Translater = new(BlepharismaMacronuclear)
var _ AltNamer = new(BlepharismaMacronuclear)
var _ IDer = new(BlepharismaMacronuclear)
var _ StartCodoner = new(BlepharismaMacronuclear)
var _ StopCodoner = new(BlepharismaMacronuclear)

type (
	// BlepharismaMacronuclear is the DNA to protein translation table
	BlepharismaMacronuclear struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s BlepharismaMacronuclear) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"TAT": 'Y', "TAC": 'Y', "TAG": 'Q', "TGT": 'C',
		"TGC": 'C', "TGG": 'W', "CTT": 'L', "CTC": 'L',
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

// String provides a human-readable indication of usage
func (s BlepharismaMacronuclear) String() string {
	return "Blepharisma Macronuclear Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s BlepharismaMacronuclear) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s BlepharismaMacronuclear) ID() uint {
	return 15
}

// StartCodons lists the codons which start a transcript
func (s BlepharismaMacronuclear) StartCodons() []string {
	return []string{"ATG"}
}

// StopCodons lists the codons which end a transcript
func (s BlepharismaMacronuclear) StopCodons() []string {
	return []string{"TAA", "TGA"}
}
