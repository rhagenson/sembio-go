# BiGr -- **B**io**I**nformatics **Gr**ammar

## Overview

`bigr` is a bioinformatics library that leverages Go's strange little way of doing things from package-scoped imports to implicit interfacing. Just like the Go is an opinionated language that does things its way, this library is opinionated.

## Order of Attribute Precedence

This library emphasizes the following in order of decreasing importance:

1.  Minimalist API Design
    -   Do not tell me anything more than I need to hear
    -   Methods should be named by what they do not how they do it
    -   Structs should be named by what they hold not how they hold it.
2.  Modularity
    -   There are often many approaches and they should be interchangeable, but well-defined in separation
    -   Scope things appropriately so each step down an import/directory tree is a further decision about what is needed
3.  Tested
    -   When I do begin accepting contributions I will require a constant increase in test coverage and example usage
4.  Program by Contract
    -   Assume everything is user input and work accordingly such as you can handle many cases, but return a consistent result
    -   In short, accept interfaces, return structs
5.  Performance
    -   Go is a compiled language so it should take advantage of being closer to the metal
    -   Clarity is more important than performance
    -   When in doubt, do it both your way and the obvious way with a test showing equivalence, benchmarks for both, and a reference to where one can further understand why the non-obvious way works
    -   I like fast code just as much as the next programmer, but being able to understand the code without unrolling all the intense thinking that went into it is better
