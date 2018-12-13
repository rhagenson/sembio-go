package fasta

const (
	// FastaHeaderPrefix is the character used by FASTA format to indicate a header line
	// Any lines NOT prefixed with FastaHeaderPrefix are considered body/sequence lines.
	FastaHeaderPrefix = '>'
)
