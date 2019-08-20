package fastq

import (
	"github.com/bio-ext/bio-go/bio/io/fasta"
)

var _ fasta.Interface = *new(Interface)

// Interface is the basic functionality of FASTQ format
// FASTQ is a four-line format that extends FASTA.
type Interface interface {
	// Header is the header line (may be internally delimited)
	Header() string

	// Sequence is the sequence line with newlines removed
	Sequence() string

	// Qualuty is the quality encoding line
	Quality() string
}
