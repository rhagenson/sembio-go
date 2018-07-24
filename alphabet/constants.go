package alphabet

// The general collection of different single byte biological alphabets
var (
	// Dna is the four letter standard encoding
	Dna = New(DnaLetters, 1)
	// DnaIupac is the sixteen letter IUPAC encoding
	DnaIupac = New(DnaIupacLetters, 1)

	// Rna is the four letter standard encoding
	Rna = New(RnaLetters, 1)

	// RnaIupac is the sixteen letter IUPAC encoding
	RnaIupac = New(RnaLetters, 1)

	// Protein is the twenty letter standard encoding
	Protein = New(ProteinLetters, 1)

	// Protein is the twenty letter standard encoding plus a gap letter
	ProteinGapped = New(ProteinGappedLetters, 1)
)

const (
	// GapLetter is the character used to represent a sequence gap.
	GapLetter = "-"

	// IupacLetters are the IUPAC ambiguous encodings
	IupacLetters = "RYSWKM" + "BDHVN"

	// DnaLetters is the strict four-letter representation of DNA.
	//	ATGC: Any of one nucelotide codes (i.e., 4 choose 1)
	DnaLetters = "ATGC"

	// DnaIupacLetters is the IUPAC representation of DNA.
	// 	ATGC: Any of one nucelotide codes (i.e., 4 choose 1)
	//	RYSWKM: Any of two nucelotide codes (i.e., 4 choose 2)
	//	BDHVN: Any of three nucleotide codes (i.e., 4 choose 3)
	//	-: Gap code (i.e., 4 choose 0)
	DnaIupacLetters = DnaLetters + IupacLetters + GapLetter

	// RnaLetters is the strict four-letter representation of DNA.
	//	AUGC: Any of one nucelotide codes (i.e., 4 choose 1)
	RnaLetters = "AUGC"

	// RnaIupacLetters is the IUPAC representation of RNA.
	//  AUGC: Any of one nucelotide codes (i.e., 4 choose 1)
	//	RYSWKM: Any of two nucelotide codes (i.e., 4 choose 2)
	//	BDHVN: Any of three nucleotide codes (i.e., 4 choose 3)
	//	-: Gap code (i.e., 4 choose 0)
	RnaIupacLetters = RnaLetters + IupacLetters + GapLetter

	// ProteinLetters is the gapless standard protein letters.
	ProteinLetters = "ACDEFGHIKLMNPQRSTVWY"

	// ProteinGappedLetters is the gapped standard protein letters.
	ProteinGappedLetters = ProteinLetters + GapLetter
)
