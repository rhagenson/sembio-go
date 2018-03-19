// Copyright 2017 Ryan Hagenson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sequence

import "bitbucket.org/rhagenson/bigr/interfaces/alphabet"

// StrictDNA is any representation that uses the StrictDNAAlphabet, which
// only has four (4) letters that are possible.
type StrictDNA interface {
	Alphabet() alphabet.DNAStrict
}

// IupacDNA is any representation that uses the IUPACDNAAlphabet, which
// has all sixteen (16) letters that are possible.
type IupacDNA interface {
	Alphabet() alphabet.DNAIupac
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
