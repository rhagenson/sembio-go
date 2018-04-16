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
	Length() uint

	// Position returns the n-th element
	Position(n uint) alphabet.Letter

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop uint) []alphabet.Letter
}

// Persistence describes methods that can be applied with a fully persistent model to create
// a new Sequence with the given new data
// This contract is should be satisfied not by returning the receiver with overwritten
// fields, but rather pointers to the old fields with a point to the change field
type Persistence interface {
	// WithPositionAs mutates a single position to the given letter, returning a new object
	WithPositionAs(n uint, letter alphabet.Letter) *Persistence

	// WithRangeAs mutates a given range of positions to the given letter array, returning a new object
	WithRangeAs(start, stop uint, letters []alphabet.Letter) *Persistence
}
