---
layout: default
title: Welcome
nav_order: 1
description: "bio-go is a semantic Bioinformatics library for the Go programming language."
permalink: /
---

# Welcome

The intent of this documentation is to introduce you to the ideas underlying `bio-go` so that you can jump into the codebase and start building useful tools.
`bio-go` is intended to unify the presentation of different approaches to the same Bioinformatics problems to allow a flexibility in design.

## Philosophy

There is hardly ever only one way to solve a Bioinformatics problem. By providing a clear semantic structure to the codebase, clarifying expectations from code at each level down the import path, `bio-go` addresses the problem of aggregating potential solutions. It should be the case that you can build a tool today and switch out parts of the it with other approaches tomorrow.

### Statement of Need

Bioinformatics projects often require building custom, small tools for the purpose of complementing larger tools' inflexibility. Developers are forced to balance size with flexibility -- the larger the program and the more people using it, the less flexible its developers can be to changing it for the better. `bio-go` is intended to be used in producing flexible tools that can change according to the specific needs of ongoing Bioinformatics projects.

`bio-go` was designed to be approachable by the end-user programmers who make up a large portion of practicing Bioinformaticians -- those with immediate research problems to solve, but without the time or resources to build and thoroughly test a multitude of solutions to the same intermediate problems. The programmer should explore solutions laterally across the import tree (e.g., `bio/sequence/immutable` to `bio/sequence/mutable`) until the approach is tuned to the specific needs of the current research problem.

## Approach

`bio-go` uses three levels of semantics to define a solution:

1. _Why_
2. _How_
3. _What_

_Why_ is the reason you are looking at using an open-source library, _How_ is the switchable approaches you have to choose from, and _What_ is the concrete elements you need to solve your problem.

### IUPAC DNA Example

For example, if we build a tool that needs to represent `100` IUPAC-encoded DNA sequences then think of the three questions:

1. Why are you looking into this module? (i.e., the kind of work you are doing)
2. How are you hoping to get the job done? (i.e., do you want speed, immutability, simplicity, and so on)
3. What are you going to use? (i.e., the "what" that allows the "why" for your work)

For our scenario of needing represent `100` IUPAC-encoded DNA sequence we might first decide:

1. _Why_: Because the tool needs to do this
2. _How_: We do not care
3. _What_: Obviously, IUPAC DNA sequences

Notice that those initial answers do not tell us much about our solution so practically any solution should work equally well.
Therefore, we could then start out by using code from `bio/sequence/immutable` to represent the sequences as `immutable.DnaIupac` objects.
Now let's now say we will want to __mutate__ the `100` sequences -- perhaps because the immutable data is slow or because we want to build a null model based on random mutation.
In order to mutate our sequences we need to make a simple modification: instead of using `bio/sequence/immutable` we instead use `bio/sequence/mutable`.
Modify our code to not use `immutable.DnaIupac` objects, but instead `mutable.DnaIupac` objects.
Why place these two in separate packages? Why not have a `MutableSeq` and `ImmutableSeq` similar to [Biopython](https://biopython.org/docs/1.74/api/Bio.Seq.html)?
Due to the way Go handles imports, the final part of the import path becomes part of each use of the object (i.e., `immutable.DnaIupac` is replaced by `mutable.DnaIupac`).
As well, we can clearly see from the import statement if either is being used, or even if both are being used.
Furthermore, we can switch between the two by changing the import line then letting the compiler tell us every location in the code that needs updating.
