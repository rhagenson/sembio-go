// Copyright 2017 Ryan Hagenson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sequence

import "bitbucket.org/rhagenson/bigr/interfaces/alphabet"

// StrictProtein is any representation that uses the ProteinStrict Alphabet
type StrictProtein interface {
	Alphabet() alphabet.ProteinStrict
}

// GappedProtein is any representation that uses the ProteinGapped Alphabet
type GappedProtein interface {
	Alphabet() alphabet.ProteinGapped
}

// StrictProteinSequence is the combination of implementing both
// StrictProtein and Sequence interfaces
type StrictProteinSequence interface {
	StrictProtein
	Sequence
}

// GappedProteinSequence is the combination of implementing both
// GappedProtein and Sequence interfaces
type GappedProteinSequence interface {
	GappedProtein
	Sequence
}
