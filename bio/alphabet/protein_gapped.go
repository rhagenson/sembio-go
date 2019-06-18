package alphabet

// ProteinGapped is the twenty letter standard encoding plus a gap letter
type ProteinGapped struct {
	*Struct
}

func NewProteinGapped() *ProteinGapped {
	return &ProteinGapped{
		New(ProteinLetters + GapLetter),
	}
}
