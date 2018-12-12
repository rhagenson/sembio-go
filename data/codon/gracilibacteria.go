package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(Gracilibacteria)
var _ Translater = new(Gracilibacteria)
var _ AltNamer = new(Gracilibacteria)
var _ IDer = new(Gracilibacteria)
var _ StartCodoner = new(Gracilibacteria)
var _ StopCodoner = new(Gracilibacteria)

type (
	// Gracilibacteria is the NCBI DNA to protein translation table
	// for gracilibacteria
	Gracilibacteria struct{}

	// CandidateDivisionSR1 is the mtDNA to protein translation table
	// for Candidate Division SR1
	CandidateDivisionSR1 Gracilibacteria
)

// Translate converts a codon into its amino acid equivalent
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

// String provides a human-readable indication of usag
func (s Gracilibacteria) String() string {
	return "Gracilibacteria Codon Library"
}

// String provides a human-readable indication of usag
func (s CandidateDivisionSR1) String() string {
	return "Candidate Division SR1 Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s Gracilibacteria) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s Gracilibacteria) ID() uint {
	return 25
}

// StartCodons lists the codons which start a transcript
func (s Gracilibacteria) StartCodons() []string {
	return []string{"TTG", "ATG", "GTG"}
}

// StopCodons lists the codons which end a transcript
func (s Gracilibacteria) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
