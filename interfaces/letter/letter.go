package letter

// Letter is an arbitrary entry in an Alphabet (string allows multirune entries in the Alphabet)
type Letter string

func (l Letter) String() string { return string(l) }
