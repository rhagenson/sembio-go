package codon

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
