package interfaces

// The unchanging alphabets of the two standard DNA considerations:
// 1) The strict DNA alphabet with only the unambiguous DNA letters, and
// 2) The ambiguous IUPAC DNA alphabet that covers all combinations of
// possibilities (see https://www.bioinformatics.org/sms/iupac.html
// for meaning of each letter).
//
// A result of these representations is that, strictly speaking, DNA is
// unambiguous and gapless so the IUPAC representation is needed
// whenever ambiguity or gaps are possible.
type StrictDNAAlphabetType [4]string
type IupacDNAAlphabetType [16]string

var (
	StrictDNAAlphabet = StrictDNAAlphabetType([4]string{"A", "T", "G", "C"})
	IupacDNAAlphabet  = IupacDNAAlphabetType([16]string{
		"A", "T", "G", "C", // Any of one nucelotide codes (i.e., 4 choose 1)
		"R", "Y", "S", "W", "K", "M", // Any of two nucelotide codes (i.e., 4 choose 2)
		"B", "D", "H", "V", "N", // Any of three nucleotide codes (i.e., 4 choose 3)
		"-", // Gap code (i.e., 4 choose 0)
	})
)

// StrictDNA is any representation that uses the StrictDNAAlphabet, which
// only has four (4) letters that are possible.
type StrictDNA interface {
	AlphabetType() StrictDNAAlphabetType
}

// IupacDNA is any representation that uses the IUPACDNAAlphabet, which
// has all sixteen (16) letters that are possible.
type IupacDNA interface {
	AlphabetType() IupacDNAAlphabetType
}

// StrictDNASequence is the combination of implementing both
// StrictDNA and Sequence interfaces
type StrictDNASequence interface {
	StrictDNA
	Sequence
}

// IupacDNASequence is the combination of implementing both
// IupacDNA and Sequence interfaces
type IupacDNASequence interface {
	IupacDNA
	Sequence
}
