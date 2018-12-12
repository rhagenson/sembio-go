package codon

import (
	"fmt"
	"strings"
)

var _ fmt.Stringer = new(SGC0)
var _ Translater = new(SGC0)
var _ AltNamer = new(SGC0)
var _ StartCodoner = new(SGC0)
var _ StopCodoner = new(SGC0)

type SGC0 struct{}
type Standard SGC0

func (s SGC0) Translate(c string) (byte, bool) {
	aa, ok := map[string]byte{
		"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L',
		"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C',
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
		"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
		"GGG": 'G',
	}[c]
	return aa, ok
}

func (s SGC0) String() string {
	return "SGC0 Codon Library"
}

func (s Standard) String() string {
	return "Standard Codon Library"
}

func (s SGC0) AltName() string {
	return "SGC0"
}

func (s SGC0) StartCodons() []string {
	return []string{"TTG", "CTG", "ATG"}
}

func (s SGC0) StopCodons() []string {
	return []string{"TTA", "TAG", "TGA"}
}

func translateDna(b string) byte {
	switch {
	case strings.HasPrefix(b, "GC"): // GCT, GCC, GCA, GCG (GCN)
		return 'A' // Alanine / Ala
	case strings.HasPrefix(b, "CG") || b == "AGA" || b == "AGG": // CGT, CGC, CGA, CGG, AGA, AGG (CGN, MGR)
		return 'R' // Arginine / Arg
	case strings.HasPrefix(b, "AA") && (b[2] == 'T' || b[2] == 'C'): // AAT, AAC (AAY)
		return 'N' // Asparagine / Asn
	case strings.HasPrefix(b, "GA") && (b[2] == 'T' || b[2] == 'C'): // GAT, GAC (GAY)
		return 'D' // Aspartic acid / Asp
	case strings.HasPrefix(b, "TG") && (b[2] == 'T' || b[2] == 'C'): // TGT, TGC (TGY)
		return 'C' // Cysteine / Cys
	case strings.HasPrefix(b, "CA") && (b[2] == 'A' || b[2] == 'G'): // CAA, CAG (CAR)
		return 'Q' // Glutamine / Gln
	case strings.HasPrefix(b, "GA") && (b[2] == 'A' || b[2] == 'G'): // GAA, GAG (GAR)
		return 'E' // Glutamic acid / Glu
	case strings.HasPrefix(b, "GG"): // GGT, GGC, GGA, GGG (GGN)
		return 'G' // Glycine / Gly
	case strings.HasPrefix(b, "CA") && (b[2] == 'T' || b[2] == 'C'): // CAT, CAC (CAY)
		return 'H' // Histidine / His
	case strings.HasPrefix(b, "AT") && (b[2] == 'T' || b[2] == 'C' || b[2] == 'A'): // ATT, ATC, ATA  (ATH)
		return 'I' // Isoleucine / Ile
	case strings.HasPrefix(b, "CT") || strings.HasPrefix(b, "TT") && (b[2] == 'A' || b[2] == 'G'): // TTA, TTG, CTT, CTC, CTA, CTG (YTR, CTN)
		return 'L' // Leucine / Leu
	case strings.HasPrefix(b, "AA") && (b[2] == 'A' || b[2] == 'G'): // AAA, AAG (AAR)
		return 'K' // Lysine / Lys
	case b == "ATG": // ATG (ATG)
		return 'M' // Start / Methionine / Met
	case strings.HasPrefix(b, "TT") && (b[2] == 'T' || b[2] == 'C'): // TTT, TTC (TTR)
		return 'F' // Phenylalanine / Phe
	case strings.HasPrefix(b, "CC"): // CCT, CCC, CCA, CCG (CCN)
		return 'P' // Proline / Pro
	case strings.HasPrefix(b, "TC") || strings.HasPrefix(b, "AG") && (b[2] == 'T' || b[2] == 'C'): // TCT, TCC, TCA, TCG, AGT, AGC (TCN, AGY)
		return 'S' // Serine / Ser
	case strings.HasPrefix(b, "AC"): // ACT, ACC, ACA, ACG (ACN)
		return 'T' // Threonine / Thr
	case b == "TGG": // TGG (TGG)
		return 'W' // Tryptophan / Tyr
	case strings.HasPrefix(b, "TA") && (b[2] == 'T' || b[2] == 'C'): // TAT, TAC (TAY)
		return 'Y' // Tyrosine / Tyr
	case strings.HasPrefix(b, "GT"): // GTT, GTC, GTA, GTG (GTN)
		return 'V' // Valine / Val
	case b == "TAA", b == "TGA", b == "TAG": // TAA, TGA, TAG (TAR, TRA)
		return '-' // Stop / Ochre / Opal / Amber
	default:
		return 'X'
	}
}
