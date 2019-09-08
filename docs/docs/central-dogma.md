---
layout: page
title:  "Example: Central Dogma"
nav_order: 2
---

# Example: Central Dogma

On this page we will explore the following series of steps:

1. Read in a given FASTA file with DNA sequences
2. Translate each sequence to protein sequences
3. Writing the protein sequences to a new FASTA file

## Initial File

We will build this out as a command-line interface (CLI) tool. It will take two CLI options:

+ `-input`, name of the input FASTA
+ `-output`, name of the output FASTA

Although there are better CLI builders out there, we will use the standard library's `flag` package.

```go
package main

import (
    "flag"
    "os"
    "fmt"
)

func main() {
    input := flag.String("input", "", "The input FASTA file (DNA)")
    output := flag.String("output", "", "The output FASTA file (Protein)")
    flag.Parse()

    if *input == "" || *output == "" {
        fmt.Fprintf(os.Stderr, "Both -input and -output are required.\n")
        flag.Usage()
        os.Exit(1)
    }
}
```

The above file is just the start. It does nothing more than build the interface between the tool and the outside world by providing an `-input` flag, `-output` flag, and an error if neither is provided.

## Reading in FASTA with DNA

```go
    x, err := os.Open(input)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Could not read the input file: %s\n", err)
        os.Exit(2)
    }
    f, err := base.ReadDna(x)
```

## Translating Sequences

## Writing out FASTA with Protein
