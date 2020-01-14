---
title: '`sembio/go`: Semantic Bioinformatics for Go'
tags:
  - 
authors:
 - name: Ryan A. Hagenson
   orcid: 0000-0001-9750-1925
   affiliation: "1"
affiliations:
 - name: Omaha's Henry Doorly Zoo and Aquarium
   index: 1
date: 15 June 2019
bibliography: paper.bib
---

# Summary

`sembio/go` starts from the known intersection of Bioinformatics data types as abstract data types (interfaces), providing functions abstracted over these abstract data types, as well as concrete data types implementing these interfaces to bootstrap development. Semantics of abstract data types are validated via extensive property testing. Flexibility is achieved through a separation of "why-how-what" at the import level such that _why_ one is using the library does not force one into _how_ it is done or _what_ the exact details of the implementation become.

Bioinformatics libraries that informed the development of this one are: BioPython [@BioPython], Rust-Bio [@Rust-Bio], and Bioconductor [@Bioconductor]. Property testing is done via <https://github.com/leanovate/gopter/>.

`sembio/go` is intended to clarify the intersection of existing solutions and define the minimal interface of biological data types such that we may abstract out the semantics of our solutions. Meanwhile, projects like `biogo` [@BioGo] are intended to set out a unified toolkit that makes definitive decisions about the semantics that `sembio/go` leaves abstract. In short, the unified approach of `biogo` is one solution possibly exported by `sembio/go`, meanwhile the reverse would not be expected.

# Statement of Need

Bioinformatics projects often require building custom, small tools for the purpose of complementing larger tools' inflexibility. Developers are forced to balance size with flexibility -- the larger the program and the more people using it, the less flexible its developers can be to changing it for the better. `sembio/go` is intended to be used in producing flexible tools that can change according to the specific needs of ongoing Bioinformatics projects.

`sembio/go` was designed to be approachable by the end-user programmers who make up a large portion of practicing Bioinformaticians -- those with immediate research problems to solve, but without the time or resources to build and thoroughly test a multitude of solutions to the same intermediate problems. The programmer should explore solutions laterally across the import tree (e.g., `bio/sequence/immutable` to `bio/sequence/mutable`) until the approach is tuned to the specific needs of the current research problem.

# Design Structure

`sembio/go` is structured intentionally to form a fat tree where each step down the tree asks/answers a new question in the following order:

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

# Acknowledgements

I acknowledge the valuable input on Bioinformatics data types and their semantics afforded to me by Sean C. West and Ryan P. Ehrlich, the friendly ears of Cynthia L. Frasier and Timothy M. Sefczek who have helped me determine how to develop with biologists as end users, and the two reviewers Will Rowe and Dan Kortschak who identified inadequacies in the initial work.

# References
