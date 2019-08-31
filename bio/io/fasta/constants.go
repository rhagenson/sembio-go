package fasta

const (
	// HeaderPrefix is the character used by FASTA format to indicate a header line
	// Any lines NOT prefixed with HeaderPrefix are considered body/sequence lines.
	HeaderPrefix = '>'
)
