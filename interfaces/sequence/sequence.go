// Copyright 2017 Ryan Hagenson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sequence

import "bitbucket.org/rhagenson/bigr/interfaces/letter"

// Sequence is an abstract type defining the basic functionality of any biological sequence (DNA, RNA, Protein, or some other series of characters in a given alphabet)
type Sequence interface {
	// Length returns how many elements there are in the Sequence
	Length() int

	// Position returns the n-th letter.Letter element
	Position(n int) letter.Letter

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop int) []letter.Letter
}
