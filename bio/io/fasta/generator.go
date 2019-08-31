package fasta

// Generator is any func that can create a FASTA from its strings
type Generator func(head, body string) (Interface, error)
