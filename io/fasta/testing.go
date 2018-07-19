package fasta

import (
	"bytes"
	"io"
	"math/rand"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

// RandomFasta generates a random valid FASTA input
// This generator is based on the very broad definition of FASTA being
// a two-line format with header and body line(s). It generates n 80-length
// lines using the letters from alphabet.Interface
func RandomFasta(seed int64, n uint, a alphabet.Interface) io.Reader {
	rand.Seed(seed)
	valid := a.Copy()
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
			b[i] = fastaHeaderPrefix
		case i%linelen == 0:
			b[i] = '\n'
		default:
			b[i] = valid[rand.Intn(len(valid))]
		}
	}

	return bytes.NewReader(b)
}
