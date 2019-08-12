# Bio -- Bioinformatics Module

[![Build Status](https://travis-ci.org/bio-ext/bio-go.svg?branch=master)](https://travis-ci.org/bio-ext/bio-go)
[![Go Report](https://goreportcard.com/badge/github.com/bio-ext/bio-go)](https://goreportcard.com/report/github.com/bio-ext/bio-go)

## Overview

`bio` is a semantic Bioinformatics module that emphasizes a definitive design structure.

## Design Structure

This library is structured in an intentional manner to build out a fat tree of ideally max depth of three where each step down the tree asks/answers a new question in the following order:

1. _Why_ are you looking into this library? (i.e., the kind of work you are doing)
2. _How_ are you hoping to get the job done? (i.e., do you need speed, immutability, simplicity, and so on)
3. _What_ are you going to use? (i.e., the what that allows the why for your work)

These questions are very general so a contrived example might be "I need to represent N sequences of DNA with IUPAC ambiguity":

1. _Why_...? Ex: I need sequences (look in `bio/sequence` package)
2. _How_...? Ex: I prefer immutable data structures (look in `.../sequence/immutable` package)
3. _What_...? Ex: I need IUPAC DNA (look in `.../immutable/*iupac*` files)

Full path: `bio/sequence/immutable/dna_iupac.go` to use `immutable.NewDnaIupac(...)` N times.

This structure should promote quick searches for the _why_, _how_, and _what_ that must be answered for every project many times over.

This design means that everything under a directory _should_ implement the interfaces above and inline with it in the tree; for example: everything under `bio/sequence` implements `sequence.Interface` and everything under `bio/alphabet` implements `alphabet.Interface` and so on.
