package fasta

// Interface is the basic functionality of FASTA format
// FASTA is a two-line format so should be possible to return either
// the header or body line(s).
type Interface interface {
	// Header is the header line (may be internally delimited)
	Header() string

	// Sequence is the sequence line with newlines removed
	Sequence() string
}
