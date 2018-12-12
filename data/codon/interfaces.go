package codon

type Translater interface {
	Translate(string) (byte, bool)
}

type AltNamer interface {
	AltName() string
}

type StartCodoner interface {
	StartCodons() []string
}

type StopCodoner interface {
	StopCodons() []string
}
