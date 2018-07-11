package alphabet

// Gap is the character used to represent a sequence gap.
const GapLetter = "-"

// DnaStrictLetters is the strict four-letter representation of DNA.
//	ATGC: Any of one nucelotide codes (i.e., 4 choose 1)
const DnaStrictLetters = "ATGC"

// DnaIupacLetters is the IUPAC representation of DNA.
// 	ATGC: Any of one nucelotide codes (i.e., 4 choose 1)
//	RYSWKM: Any of two nucelotide codes (i.e., 4 choose 2)
//	BDHVN: Any of three nucleotide codes (i.e., 4 choose 3)
//	-: Gap code (i.e., 4 choose 0)
const DnaIupacLetters = DnaStrictLetters + "RYSWKM" + "BDHVN" + GapLetter

// rnaStrictLetters is the strict four-letter representation of DNA.
//	AUGC: Any of one nucelotide codes (i.e., 4 choose 1)
const RnaStrictLetters = "AUGC"

// RnaIupacLetters is the IUPAC representation of RNA.
//  AUGC: Any of one nucelotide codes (i.e., 4 choose 1)
//	RYSWKM: Any of two nucelotide codes (i.e., 4 choose 2)
//	BDHVN: Any of three nucleotide codes (i.e., 4 choose 3)
//	-: Gap code (i.e., 4 choose 0)
const RnaIupacLetters = RnaStrictLetters + "RYSWKM" + "BDHVN" + GapLetter

// ProteinGapped is the gapless standard protein letters.
const ProteinStrictLetters = "ACDEFGHIKLMNPQRSTVWY"

// ProteinGapped is the gapped standard protein letters.
const ProteinGappedLetters = ProteinStrictLetters + GapLetter
