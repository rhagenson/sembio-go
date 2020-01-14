package fastq

import "github.com/sembio/go/bio/io/fasta"

// Interface is the basic functionality of FASTQ format
// FASTQ is a four-line format that extends FASTA.
type Interface interface {
	// fasta.Interface defines: Header() and Sequence()
	fasta.Interface

	// Quality is the quality encoding line
	Quality() string
}
