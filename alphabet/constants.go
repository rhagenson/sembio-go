package alphabet

// The general collection of different single byte biological alphabets
const (
	// Dna is the four letter standard encoding
	Dna = Alphabet(DnaLetters)

	// DnaIupac is the sixteen letter IUPAC encoding
	DnaIupac = Alphabet(DnaIupacLetters)

	// Rna is the four letter standard encoding
	Rna = Alphabet(RnaLetters)

	// RnaIupac is the sixteen letter IUPAC encoding
	RnaIupac = Alphabet(RnaIupacLetters)

	// Protein is the twenty letter standard encoding
	Protein = Alphabet(ProteinLetters)

	// Protein is the twenty letter standard encoding plus a gap letter
	ProteinGapped = Alphabet(ProteinGappedLetters)
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
