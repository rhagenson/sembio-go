## Order of Attribute Priority

This library emphasizes the following in order of decreasing importance:

1. Minimalist API Design
    - _Do not tell me anything more than I need to hear. Do not listen just to ignore._
    - Methods should be named by what they do not how they do it
    - Structs should be named by what they hold not how they hold it
2. Modularity
    - There are often many approaches -- they should be interchangeable and well-defined in separation
    - Scope things appropriately so each step down an import/directory tree is a further decision about what is needed
3. Tested
    - An implementation without tests is untrusted
4. Program by Contract
    - Assume everything is user input and work accordingly such that you can handle many cases, but return a consistent result
    - Accept interfaces, return structs
5. Performance
    - Go is a compiled language so it can take advantage of being "closer to the metal"
    - Clarity is more important than performance
    - If an solution is not non-obvious, provide a reference and show equivalence with the naive solution in property test(s)
    - Being able to understand the code without unrolling all the intense thinking that went into it is better than the fastest library code
