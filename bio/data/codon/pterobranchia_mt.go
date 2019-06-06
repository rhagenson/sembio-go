package codon

import (
	"fmt"
)

var _ fmt.Stringer = new(PterobranchiaMt)
var _ Translater = new(PterobranchiaMt)
var _ AltNamer = new(PterobranchiaMt)
var _ IDer = new(PterobranchiaMt)
var _ StartCodoner = new(PterobranchiaMt)
var _ StopCodoner = new(PterobranchiaMt)

type (
	// PterobranchiaMt is the pterobranchia mtDNA to protein translation tabl
	PterobranchiaMt struct{}
)

// Translate converts a codon into its amino acid equivalent
func (s PterobranchiaMt) Translate(c string) (byte, bool) {
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
		"AGA": 'S', "AGG": 'K', "GTT": 'V', "GTC": 'V',
		"GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A',
		"GCA": 'A', "GCG": 'A', "GAT": 'D', "GAC": 'D',
		"GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G',
		"GGA": 'G', "GGG": 'G',
	}[c]
	return aa, ok
}

// String provides a human-readable indication of usag
func (s PterobranchiaMt) String() string {
	return "Pterobranchia Mitochondrial Codon Library"
}

// AltName provides the alternative name used by NCBI
func (s PterobranchiaMt) AltName() string {
	return ""
}

// ID provides the alternative identifier used by NCBI
func (s PterobranchiaMt) ID() uint {
	return 24
}

// StartCodons lists the codons which start a transcript
func (s PterobranchiaMt) StartCodons() []string {
	return []string{"TTG", "CTG", "ATG", "GTG"}
}

// StopCodons lists the codons which end a transcript
func (s PterobranchiaMt) StopCodons() []string {
	return []string{"TAA", "TAG"}
}
