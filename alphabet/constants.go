package alphabet

// GapLetter is the character used to represent a sequence gap.
const GapLetter = "-"

// DnaLetters is the strict four-letter representation of DNA.
//	ATGC: Any of one nucelotide codes (i.e., 4 choose 1)
const DnaLetters = "ATGC"

// DnaIupacLetters is the IUPAC representation of DNA.
// 	ATGC: Any of one nucelotide codes (i.e., 4 choose 1)
//	RYSWKM: Any of two nucelotide codes (i.e., 4 choose 2)
//	BDHVN: Any of three nucleotide codes (i.e., 4 choose 3)
//	-: Gap code (i.e., 4 choose 0)
const DnaIupacLetters = DnaLetters + "RYSWKM" + "BDHVN" + GapLetter

// RnaLetters is the strict four-letter representation of DNA.
//	AUGC: Any of one nucelotide codes (i.e., 4 choose 1)
const RnaLetters = "AUGC"

// RnaIupacLetters is the IUPAC representation of RNA.
//  AUGC: Any of one nucelotide codes (i.e., 4 choose 1)
//	RYSWKM: Any of two nucelotide codes (i.e., 4 choose 2)
//	BDHVN: Any of three nucleotide codes (i.e., 4 choose 3)
//	-: Gap code (i.e., 4 choose 0)
const RnaIupacLetters = RnaLetters + "RYSWKM" + "BDHVN" + GapLetter

// ProteinLetters is the gapless standard protein letters.
const ProteinLetters = "ACDEFGHIKLMNPQRSTVWY"

// ProteinGappedLetters is the gapped standard protein letters.
const ProteinGappedLetters = ProteinLetters + GapLetter
