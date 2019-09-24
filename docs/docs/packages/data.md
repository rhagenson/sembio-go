---
layout: page
title:  "Data"
nav_order: 2
heading_anchors: true
parent: Packages
---

## Data

There are some recurring pieces of data we Bioinformatians tend to need codified.
This package contains those pieces of data.

### codon

This package contains the 31 DNA-to-Protein codon lookup tables from <https://www.ncbi.nlm.nih.gov/Taxonomy/Utils/wprintgc.cgi>.

There are a few interfaces related to codon tables:

```go
// Interface is a full featured codon lookup table
type Interface interface {
    Translater
    AltNamer
    IDer
    StartCodoner
    StopCodoner
}

// Translater converts a codon into its amino acid equivalent
type Translater interface {
    Translate(string) (byte, bool)
}

// AltNamer provides the alternative name for a codon translation table
// Zero value ("") indicates no alternative name.
type AltNamer interface {
    AltName() string
}

// IDer provides the ID code used by NCBI
type IDer interface {
    ID() uint
}

// StartCodoner lists the codons which start a transcript
type StartCodoner interface {
    StartCodons() []string
}

// StopCodoner lists the codons which end a transcript
type StopCodoner interface {
    StopCodons() []string
}
```

A complete codon lookup table satisfies `Interface`.
A codon lookup table that has no alternative name produces an empty string.
`Translate(string) (byte, bool)` is the corresponding amino acid code and no error, or no character and an error if the codon was not found.
