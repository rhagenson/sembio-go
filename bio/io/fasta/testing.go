package fasta

import (
	"math/rand"

	"github.com/bio-ext/bio-go/bio/alphabet"
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
func TestGenMultiFasta(seed int64, n, m uint, a alphabet.Interface) []byte {
	rand.Seed(seed)
	nseqs := rand.Intn(int(m))
	b := make([]byte, 0)
	for i := 0; i < nseqs; i++ {
		b = append(b, TestGenFasta(seed, n, a)...)
	}
	return b
}
