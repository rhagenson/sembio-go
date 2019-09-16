---
layout: page
title:  "Sequence"
nav_order: 2
heading_anchors: true
parent: Packages
---

## Sequence

All biological sequences have a finite, ordered series of positions.

```go
type Interface interface {
	// Length is the number of elements in the Interface
	Length() uint

	// Position is the n-th element
	Position(n uint) (string, error)

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop uint) (string, error)
}
```

In other words, a biological sequence has a definite, positive length (`Length() unit`) and can return any given position (`Position(n uint) (string, error)`) or range of positions (`Range(start, stop uint) (string, error`).

Practically speaking, there do exist more interfaces to biological sequences such as a sequence which is: reversible, can be complemented, translated, or transcribed.
Therefore the following interfaces are also defined:

```go
// Reverser can reverse the sequence
type Reverser interface {
	Reverse() (Interface, error)
}

// Complementer can complement the sequence
type Complementer interface {
	Complement() (Interface, error)
}

// RevComper can reverse-complement the sequence
type RevComper interface {
	RevComp() (Interface, error)
}

// Translater can translate the sequence
type Translater interface {
	Translate(codon.Interface, byte) (Interface, error)
}

// Transcriber can transcribe the sequence
type Transcriber interface {
	Transcribe() (Interface, error)
}

// Alphabeter s you what alphabet.Interface it uses
type Alphabeter interface {
	Alphabet() alphabet.Interface
}

// LetterCounter counts the letters observed when reading the sequence
type LetterCounter interface {
	LetterCount() map[string]uint
}
```

The opaque interfaces among these which do not seemingly explain themselves are the final two: `Alphabet() alphabet.Interface` and `LetterCount() map[string]uint`.
The former of these exists to define a way to obtain the alphabet of valid characters of a sequence (which by chance might not contain all possible valid characters), while the second exists to define a way to count the frequency of valid characters.
**Warning**: These two should not be intertwined so `LetterCount()` might not include all valid characters in its count if the sequence does not contain any instances of a certain character.
Optionally, implementations can initialize the counts of all valid characters to zero, but this should not be relied upon in the design of systems as not all concrete sequence types might handle counting letters this way.

### immutable

This version of sequence constructs "immutable" sequence which, once constructed, cannot be changed.
In actuality, sequences constructed via this package can be thought of as somewhere between "immutable" and "persistent" such that a sequence can be modified using `WithFunc`s (a type alias to `func(*Struct) *Struct`) -- an individual sequence that you might have assigned to a variable will never change (immutable), but you can mutate the sequence to construct new sequences and any intermediate sequences assigned to variables will remain (persistent).

**Point of advice**: If you plan to make use of Go's concurrency you likely want to use sequences from this package as they are safe to "mutate" concurrently (despite what the race detector states) as the starting sequence is only ever read after constructing.

### mutable

This version of sequence constructs mutable sequences, which can be modified at any time.
Mutable sequences do not create as much garbage when modifying them, but at the cost of not being safe to use between threads/goroutines.

**Word of caution**: If you are aware of what you are doing, and run the race detector to confirm there are no known data races in your solution, mutable sequences are a lot cheaper to use as there is less garbage to collect during execution.
However, concurrency is a difficult problem to get right so be careful.
