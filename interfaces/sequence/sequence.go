// Copyright 2017 Ryan Hagenson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sequence

import (
	"bitbucket.org/rhagenson/bigr/interfaces/alphabet"
)

// Sequence is an abstract type defining the basic functionality of any biological sequence
// (DNA, RNA, Protein, or some other series of Letters in a given Alphabet)
// A Sequence is essentially an ordered list of elements where there is a Lenth(), each element has a Position(),
// it is possible to look at a specific Range() of elements
type Sequence interface {
	// Length returns how many elements there are in the Sequence
	Length() int

	// Position returns the n-th element
	Position(n int) alphabet.Letter

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop int) []alphabet.Letter
}
