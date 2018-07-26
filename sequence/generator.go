package sequence

// Generator is any func that can create a sequence from its string
type Generator func(string) (Interface, error)
