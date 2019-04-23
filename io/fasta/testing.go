package fasta

import (
	"math/rand"

	"bitbucket.org/rhagenson/bio/alphabet"
)

// TestGenFasta generates a random valid FASTA input
// This generator is based on the very broad definition of FASTA being
// a two-line format with header and body line(s). It generates n 80-length
// lines using the letters from alphabet.Interface
func TestGenFasta(seed int64, n uint, a alphabet.Interface) []byte {
	rand.Seed(seed)
	valid := a.String()
	linelen := func() uint {
		if n < 80 {
			return n / 2
		}
		return 80
	}()
	tot := n * linelen
	b := make([]byte, tot)
	for i := uint(0); i < tot; i++ {
		switch {
		case i == 0:
			b[i] = FastaHeaderPrefix
		case i%linelen == 0:
			b[i] = '\n'
		default:
			b[i] = valid[rand.Intn(len(valid))]
		}
	}
	return b
}

// TestGenMultiFasta generates a random valid multiple-record FASTA input
// This generator is based on the very broad definition of FASTA being
// a two-line format with header and body line(s). It generates n 80-length
// lines using the letters from alphabet.Interface
func TestGenMultiFasta(seed int64, n uint, a alphabet.Interface) []byte {
	seqs := make([]byte, 10*int(n))
	for i := 0; i < 10*int(n); i += int(n) {
		seq := TestGenFasta(seed, n, a)
		for j, b := range seq {
			seqs[i+j] = b
		}
	}
	return seqs
}
