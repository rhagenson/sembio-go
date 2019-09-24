---
layout: page
title:  "Central Dogma"
nav_order: 2
heading_anchors: true
parent: Example
---

## Central Dogma

On this page we will explore the following series of steps:

1. TOC
{:toc}

### Set up Command Line Interface

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

The above file is just the start. It does nothing more than build the CLI between the tool and the outside world by providing an `-input` flag, `-output` flag, and an error if neither is provided.

Running this without any input results in:

```bash
Both -input and -output are required.
Usage of <name of executable>:
  -input string
        The input FASTA file (DNA)
  -output string
        The output FASTA file (Protein)
```

### Reading in FASTA with DNA

Now that we have the input and output setup and the error case of neither option being provided, we can proceed to reading in the input FASTA (which should contain DNA sequences). Note that `(...)` indicates truncated content as we needed to add an import above the previously written content and more code below the previously written content.

```go
import (
    "flag"
    "fmt"
    "os"

    fa "github.com/bio-ext/bio-go/bio/io/fasta/base"
)

(...)

    x, err := os.Open(*input)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Could not open the input file: %s\n", err)
        os.Exit(2)
    }
    f, err := fa.ReadMultiDna(x)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error in reading %s: %s\n", *input, err)
        os.Exit(2)
    }
```

The careful eye will note two decisions we have made:

1. Using an alias for our import of FASTA readers/writers (`fa "github.com/bio-ext/bio-go/bio/io/fasta/base"`)
2. We run `fa.ReadMultiDna(x)`, not `fa.ReadDna(x)`

This first decision is so that, if in the future a new FASTA reader/writer approach is added, we can change the import and the rest of the code will already refer to the `fa` identifier rather than `base` as it would be without the alias.
The second decision is because `fa.ReadDna(x)` would only give us a single FASTA record (the first one to be precise).
In `bio-go`, a distinction is made between reading a record and reading all records.
(In your own code you will likely want `fa.ReadMultiDna(x)` as we have used here. The reason for this distinction is that by definition a FASTA file _should_ only have one record, but _most_ FASTA files used today have multiple records. Blame Ryan A. Hagenson for being such a stick in the mud about this point.)

### Translating Sequences

The next step is translating the incoming FASTA records to protein. First we will allocate memory to hold the translated sequences -- we do this by making a long enough `fasta.Interface` array. Then we generate an immutable version of the underlying DNA and run the `Translate(...)` method on it with the corresponding codon table and a character of our choice to denote if a stop codon was found (here we use `~` because it is stands out visually). **Important**: Notice that we then provide a second check during the usual `err != nil` statement -- this second statement ignores errors pertaining to finding our chosen stop codon character (`~`) which is, of course, not in the standard protein alphabet.

```go
    out := make([]fasta.Interface, len(f))
    table := new(codon.Standard)
    for i, r := range f {
        d, _ := immutable.NewDna(r.Sequence())
        seq, err := d.Translate(table, '~')
        if err != nil && !strings.Contains(err.Error(), "~") {
            fmt.Fprintf(os.Stderr, "Error in translating sequence: %s\n%s\n", err, d.String())
            os.Exit(2)
        }
        out[i] = fasta.Interface(fa.New(r.Header(), seq))
    }
```

### Writing out FASTA with Protein

Lastly, we write out the protein sequences to a FASTA based on the output name provided.

```go
    o, err := os.Create(*output)
    defer o.Close()
    for i := range out {
        o.WriteString(out[i].Header() + "\n" + out[i].Sequence() + "\n")
    }
```

## Full Solution

```go
package main

import (
    "flag"
	"fmt"
	"os"
	"strings"

	"github.com/bio-ext/bio-go/bio/data/codon"
	"github.com/bio-ext/bio-go/bio/io/fasta"
	fa "github.com/bio-ext/bio-go/bio/io/fasta/base"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
)

func main() {
	// Set up CLI
	input := flag.String("input", "", "The input FASTA file (DNA)")
	output := flag.String("output", "", "The output FASTA file (Protein)")
	flag.Parse()

	if *input == "" || *output == "" {
		fmt.Fprintf(os.Stderr, "Both -input and -output are required.\n")
		flag.Usage()
		os.Exit(1)
	}

	// Reading in FASTA with DNA
	x, err := os.Open(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open the input file: %s\n", err)
		os.Exit(2)
	}
	f, err := fa.ReadMultiDna(x)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in reading %s: %s\n", *input, err)
		os.Exit(2)
	}

	// Translating Sequences
	out := make([]fasta.Interface, len(f))
	table := new(codon.Standard)
	for i, r := range f {
		d, _ := immutable.NewDna(r.Sequence())
		seq, err := d.Translate(table, '~')
		if err != nil && !strings.Contains(err.Error(), "~") {
			fmt.Fprintf(os.Stderr, "Error in translating sequence: %s\n%s\n", err, d.String())
			os.Exit(2)
		}
		out[i] = fasta.Interface(fa.New(r.Header(), seq))
	}

	// Writing out FASTA with Protein
	o, err := os.Create(*output)
	defer o.Close()
	for i := range out {
		o.WriteString(out[i].Header() + "\n" + out[i].Sequence() + "\n")
	}
}
```
