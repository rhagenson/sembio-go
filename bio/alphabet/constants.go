package alphabet

const (
	// GapLetter is the character used to represent a sequence gap.
	GapLetter = "-"

	// IupacLetters are the IUPAC ambiguous encodings
	IupacLetters = "RYSWKM" + "BDHV" + "N"

	// DnaLetters is the strict four-letter representation of DNA.
	//	ATGC: Any of one nucleotide codes (i.e., 4 choose 1)
	DnaLetters = "ATGC"

	// DnaIupacLetters is the IUPAC representation of DNA.
	//  -: Any of zero nucleotide codes (i.e., 4 choose 0)
	// 	ATGC: Any of one nucleotide codes (i.e., 4 choose 1)
	//	RYSWKM: Any of two nucleotide codes (i.e., 4 choose 2)
	//	BDHV: Any of three nucleotide codes (i.e., 4 choose 3)
	//  N: Any of four nucleotide codes (i.e., 4 choose 4)
	DnaIupacLetters = DnaLetters + IupacLetters + GapLetter

	// RnaLetters is the strict four-letter representation of DNA.
	//	AUGC: Any of one nucleotide codes (i.e., 4 choose 1)
	RnaLetters = "AUGC"

	// RnaIupacLetters is the IUPAC representation of RNA.
	//  -: Any of zero nucleotide codes (i.e., 4 choose 0)
	//  AUGC: Any of one nucleotide codes (i.e., 4 choose 1)
	//	RYSWKM: Any of two nucleotide codes (i.e., 4 choose 2)
	//	BDHV: Any of three nucleotide codes (i.e., 4 choose 3)
	//  N: Any of four nucleotide codes (i.e., 4 choose 4)
	RnaIupacLetters = RnaLetters + IupacLetters + GapLetter

	// ProteinLetters is the gapless standard protein letters.
	ProteinLetters = "ACDEFGHIKLMNPQRSTVWY"
)
