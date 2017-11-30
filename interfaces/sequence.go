// Copyright 2017 Ryan Hagenson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interfaces

// Sequence is an abstract type defining the basic functionality of any biological sequence (DNA, RNA, Protein, or some other series of characters in a given alphabet)
// string type is returned to implictly allow non-utf8, multi-character elements which would allow defining a k-mer Sequence.
type Sequence interface {
	// Length returns how many elements there are in the Sequence
	Length() int

	// Position returns the n-th element
	Position(n int) string

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop int) []string
}
