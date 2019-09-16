---
layout: page
title:  "FASTQ"
nav_order: 2
heading_anchors: true
parent: IO
---

## FASTQ

A FASTQ formatted file contains a finite number of lines and four line types:

1. header
2. sequence
3. header duplicate (optional)
4. quality

In Go, this translates into the following interface:

```go
type Interface interface {
	// fasta.Interface defines: Header() and Sequence()
	fasta.Interface

	// Quality is the quality encoding line
	Quality() string
}
```

Each set of four lines corresponds to one triple of header, sequence, and quality with an optional duplicate.
As FASTQ is an extension of the information found in a FASTA file, this interface extends that of `fasta`.

**Word of caution**: As FASTQ extends FASTA, the standard reader (`fastq.Read`) requires you to specify how many records to read.
An error should result if you request more records than are contained in the file and requesting zero (0) records will return all records.
