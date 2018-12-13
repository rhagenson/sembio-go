# Bio -- Bioinformatics Module

## Overview

`bio` is a bioinformatics module that leverages Go way of doing things from package-scoped imports to implicit interfacing.

## Order of Attribute Precedence

This library emphasizes the following in order of decreasing importance:

1. Minimalist API Design
    - _Do not tell me anything more than I need to hear. Do not listen just to ignore._
    - Methods should be named by what they do not how they do it
    - Structs should be named by what they hold not how they hold it
2. Modularity
    - There are often many approaches -- they should be interchangeable, but well-defined in separation
    - Scope things appropriately so each step down an import/directory tree is a further decision about what is needed
3. Tested
    - A constant increase in test coverage and example usage is expected through development
4. Program by Contract
    - Assume everything is user input and work accordingly such that you can handle many cases, but return a consistent result
    - Accept interfaces, return structs
5. Performance
    - Go is a compiled language so it can take advantage of being "closer to the metal"
    - Clarity is more important than performance
    - When in doubt, do it both your non-obvious way and the obvious way with a test showing equivalence, benchmarks for both, and a reference to where one can further understand why the non-obvious way works
    - I like fast code just as much as the next programmer, but being able to understand the code without unrolling all the intense thinking that went into it is better

## Design Structure

This library is structured in an intentional manner to build out a fat tree of ideally max depth of three where each step down the tree asks/answers a new question in the following order:

1. _Why_ are you looking into this library? (i.e., the kind of work you are doing)
2. _How_ are you hoping to get the job done? (i.e., do you need speed, immutability, simplicity, and so on)
3. _What_ are you going to use? (i.e., the what that allows the why for your work)

These questions are very general so a contrived example might be "I need to represent N sequences of DNA with IUPAC ambiguity":

1. _Why_...? Ex: I need sequences (look in `bio/sequence` package)
2. _How_...? Ex: I prefer immutable data structures (look in `sequence/immutable` package)
3. _What_...? Ex: I need IUPAC DNA (look in `immutable/*iupac*` files)

Full path: `bio/sequence/immutable/dna_iupac.go` to use `immutable.NewDnaIupac(...)` N times.

This structure should promote quick searches for the _why_, _how_, and _what_ that must be answered for every project many times over.

This design means that everything under a directory _should_ implement the interfaces above and inline with it in the tree; for example: everything under `bio/sequence` implements `sequence.Interface` and everything under `bio/alphabet` implements `alphabet.Interface` and so on.
