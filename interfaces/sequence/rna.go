// Copyright 2017 Ryan Hagenson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sequence

import "bitbucket.org/rhagenson/bigr/interfaces/alphabet"

// StrictRNA is any representation that uses the StrictRNAAlphabet, which
// only has four (4) letters that are possible.
type StrictRNA interface {
	Alphabet() alphabet.RNAStrict
}

// IupacRNA is any representation that uses the IUPACRNAAlphabet, which
// has all sixteen (16) letters that are possible.
type IupacRNA interface {
	Alphabet() alphabet.RNAIupac
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
