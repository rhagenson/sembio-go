package alphabet

// gap is the character used to represent a sequence gap.
const gap = "-"

// dnaStrictLetters is the strict four-letter representation of DNA.
//	ATGC: Any of one nucelotide codes (i.e., 4 choose 1)
const dnaStrictLetters = "ATGC"

// dnaIupacLetters is the IUPAC representation of DNA.
// 	ATGC: Any of one nucelotide codes (i.e., 4 choose 1)
//	RYSWKM: Any of two nucelotide codes (i.e., 4 choose 2)
//	BDHVN: Any of three nucleotide codes (i.e., 4 choose 3)
//	-: Gap code (i.e., 4 choose 0)
const dnaIupacLetters = dnaStrictLetters + "RYSWKM" + "BDHVN" + gap

// rnaStrictLetters is the strict four-letter representation of DNA.
//	AUGC: Any of one nucelotide codes (i.e., 4 choose 1)
const rnaStrictLetters = "AUGC"

// rnaIupacLetters is the IUPAC representation of RNA.
//  AUGC: Any of one nucelotide codes (i.e., 4 choose 1)
//	RYSWKM: Any of two nucelotide codes (i.e., 4 choose 2)
//	BDHVN: Any of three nucleotide codes (i.e., 4 choose 3)
//	-: Gap code (i.e., 4 choose 0)
const rnaIupacLetters = rnaStrictLetters + "RYSWKM" + "BDHVN" + gap

// proteinGapped is the gapless standard protein letters.
const proteinStrict = "ACDEFGHIKLMNPQRSTVWY"

// proteinGapped is the gapped standard protein letters.
const proteinGapped = proteinStrict + gap
