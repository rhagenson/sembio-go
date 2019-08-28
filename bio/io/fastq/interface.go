package fastq

import "github.com/bio-ext/bio-go/bio/io/fasta"

// Interface is the basic functionality of FASTQ format
// FASTQ is a four-line format that extends FASTA.
type Interface interface {
	fasta.Interface

	// Qualuty is the quality encoding line
	Quality() string
}
