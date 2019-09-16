---
layout: page
title:  "FASTA"
nav_order: 2
heading_anchors: true
parent: IO
---

## FASTA

A FASTA formatted file contains two (or more) lines with only two line types:

1. header (only one)
2. sequence (one or more)

In Go, this translates into the following interface:

```go
type Interface interface {
	// Header is the header line (may be internally delimited)
	Header() string

	// Sequence is the sequence with newlines removed
	Sequence() string
}
```

**Word of caution**: Because, by definition, a FASTA file only has one header per file, the standard reader (`fasta.Read`) requires you to specify how many records to read.
An error should result if you request more records than are contained in the file and requesting zero (0) records will return all records.
