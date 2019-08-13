---
title: '`bio-go`: A flexible tools Bioinformatics library'
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

`bio-go` starts from the known intersection of Bioinformatics data types as abstract data types (interfaces), providing functions abstracted over these abstract data types, as well as concrete data types implementing these interfaces to bootstrap development. Semantics of abstract data types are validated via extensive property testing. Flexibility is achieved through a separation of "why-how-what" at the import level such that _why_ one is using the library does not force one into _how_ it is done or _what_ the exact details of the implementation become.

Bioinformatics libraries that informed the development of this one are: BioPython [@BioPython], Rust-Bio [@Rust-Bio], and Bioconductor [@Bioconductor]. Property testing is done via <https://github.com/leanovate/gopter/>.

# Statement of Need

Bioinformatics projects often require building custom, small tools for the purpose of complementing larger tools' inflexibility. Developers are forced to balance size with flexibility -- the larger the program and the more people using it, the less flexible its developers can be to changing it for the better. `bio-go` is intended to be used in producing flexible tools that can change according to the specific needs of ongoing Bioinformatics projects.

`bio-go` was designed to be approachable by the end-user programmers who make up a large portion of practicing Bioinformaticians -- those with immediate research problems to solve, but without the time or resources to build and thoroughly test a multitude of solutions to the same intermediate problems. The programmer should explore solutions laterally across the import tree (e.g., `bio/sequence/immutable` to `bio/sequence/mutable`) until the approach is tuned to the specific needs of the current research problem.

# Acknowledgements

I acknowledge the valuable input on Bioinformatics data types and their semantics afforded to me by Sean C. West and Ryan P. Ehrlich, as well as the friendly ears of Cynthia L. Frasier and Timothy M. Sefczek who have helped me determine how to develop with biologists as end users.

# References
