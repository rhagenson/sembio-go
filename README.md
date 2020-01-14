# Bio -- Bioinformatics Module

<!-- Dynamic badges -->
[![Build Status](https://travis-ci.org/sembio/go.svg?branch=master)](https://travis-ci.org/sembio/go)
[![codecov](https://codecov.io/gh/sembio/go/branch/master/graph/badge.svg)](https://codecov.io/gh/sembio/go)
[![GoReport](https://goreportcard.com/badge/github.com/sembio/go)](https://goreportcard.com/report/github.com/sembio/go)
[![CodeFactor](https://www.codefactor.io/repository/github/sembio/go/badge)](https://www.codefactor.io/repository/github/sembio/go)

<!-- Static badges -->
[![GoDoc](https://godoc.org/github.com/sembio/go?status.svg)](https://godoc.org/github.com/sembio/go)
[![LICENSE](https://img.shields.io/github/license/sembio/go)](LICENSE)

<!-- Citations -->
[![DOI](https://zenodo.org/badge/180650332.svg)](https://zenodo.org/badge/latestdoi/180650332)
[![status](https://joss.theoj.org/papers/7217d1ef147e05e94f323aa080d10422/status.svg)](https://joss.theoj.org/papers/7217d1ef147e05e94f323aa080d10422)

## Overview

`bio` is a semantic Bioinformatics module that emphasizes a definitive design structure.

## Installation

```bash
go get github.com/sembio/go/...
```

## Design Structure

This module is structured intentionally to form a fat tree where each step down the tree asks/answers a new question in the following order:

1. _Why_ are you looking into this module? (i.e., the kind of work you are doing)
2. _How_ are you hoping to get the job done? (i.e., do you need speed, immutability, simplicity, and so on)
3. _What_ are you going to use? (i.e., the what that allows the why for your work)

These questions are very general so a contrived example might be "I need to represent N sequences of DNA with IUPAC ambiguity":

1. _Why_...? Ex: I need sequences (look in `bio/sequence` package)
2. _How_...? Ex: I prefer immutable data structures (look in `.../sequence/immutable` package)
3. _What_...? Ex: I need IUPAC DNA (look in `.../immutable/*iupac*` files)

Full path: `bio/sequence/immutable/dna_iupac.go` to use `immutable.NewDnaIupac(...)` N times.

This structure should promote quick searches for the _why_, _how_, and _what_ that must be answered for every project many times over.

This design means that everything under a directory _should_ implement the interfaces above and inline with it in the tree; for example: everything under `bio/sequence` implements `sequence.Interface` and everything under `bio/alphabet` implements `alphabet.Interface` and so on.

If more than three levels are deemed necessary the first level will represent some generic functionality, such as in the case of `bio/io/fasta/base/fasta.go` which is housed under the generic `io` then answers our three questions `fasta` (_why_), `base` (_how_), `fasta.go` (_what_).

## Documentation

Documentation can be built through use of `godoc -http=localhost:6060` which then generates documentation accessible through a web browser at `localhost:6060`. For access to  prior to downloading `sembio/go`, visit: <https://godoc.org/github.com/sembio/go/bio>

## Testing

Tests can be run through use of `go test -v ./...` at the root of this repository. Use of the verbose `-v` flag is recommended as `sembio/go` makes extensive use of property tests, which should provide insight into what can be expected of a given implementation.

All tests, benchmarks, and examples are run continuously and can be viewed on [Travis CI](https://travis-ci.org/sembio/go).
