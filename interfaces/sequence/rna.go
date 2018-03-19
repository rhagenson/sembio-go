// Copyright 2017 Ryan Hagenson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sequence

// The unchanging alphabets of the two standard RNA considerations:
// 1) The strict RNA alphabet with only the unambiguous RNA letters, and
// 2) The ambiguous IUPAC RNA alphabet that covers all combinations of
// possibilities (see https://www.bioinformatics.org/sms/iupac.html
// for meaning of each letter).
//
// A result of these representations is that, strictly speaking, RNA is
// unambiguous and gapless so the IUPAC representation is needed
// whenever ambiguity or gaps are possible.

// StrictRNAAlphabetType is the form all standard representations of RNA
// should follow. That is the alphabet contains only four letters (AUGC).
type StrictRNAAlphabetType [4]string

// IupacRNAAlphabetType is the form all non-standard or ambiguous
// representations of RNA should follow.
// That is the alphabet that contains all 16 semi-ambiguous letters
// (ATGCRYSWKMBDHVN-).
type IupacRNAAlphabetType [16]string

var (
	// StrictRNAAlphabet is a pre-defined, correct StrictRNAAlphabetType
	StrictRNAAlphabet = StrictRNAAlphabetType([4]string{"A", "U", "G", "C"})

	// IupacRNAAlphabet is a pre-define, correct IupacRNAAlphabetType
	IupacRNAAlphabet = IupacRNAAlphabetType([16]string{
		"A", "U", "G", "C", // Any of one nucelotide codes (i.e., 4 choose 1)
		"R", "Y", "S", "W", "K", "M", // Any of two nucelotide codes (i.e., 4 choose 2)
		"B", "D", "H", "V", "N", // Any of three nucleotide codes (i.e., 4 choose 3)
		"-", // Gap code (i.e., 4 choose 0)
	})
)

// StrictRNA is any representation that uses the StrictRNAAlphabet, which
// only has four (4) letters that are possible.
type StrictRNA interface {
	Alphabet() StrictRNAAlphabetType
}

// IupacRNA is any representation that uses the IUPACRNAAlphabet, which
// has all sixteen (16) letters that are possible.
type IupacRNA interface {
	Alphabet() IupacRNAAlphabetType
}

// StrictRNASequence is the combination of implementing both
// StrictRNA and Sequence interfaces
type StrictRNASequence interface {
	StrictRNA
	Sequence
}

// IupacRNASequence is the combination of implementing both
// IupacRNA and Sequence interfaces
type IupacRNASequence interface {
	IupacRNA
	Sequence
}
