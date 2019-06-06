package fastq

import (
	"math/rand"
	"strings"

	"github.com/rhagenson/bio-go/bio/alphabet"
)

// TestGenFastq generates a random valid FASTQ
// It generates n 80-length lines using the letters from alphabet.Interface
func TestGenFastq(seed int64, n uint, a alphabet.Interface) []byte {
	rand.Seed(seed)
	valid := a.String()
	header := string(FastqHeaderPrefix) + "EAS139:136:FC706VJ:2:2104:15343:197393 1:Y:18:ATCACG"
	sequence := ""
	repeatHeader := string(FastqPreQualityHeaderPrefix) + header[1:]
	quality := ""
	minQual := int(MinPhred33)
	maxQual := int(MaxIlluminaPhred33)
	linelen := func() uint {
		if n < 80 {
			return n / 2
		}
		return 80
	}()
	b := make([]byte, linelen)
	for i := range b {
		b[i] = valid[rand.Intn(len(valid))]
	}
	sequence = string(b)

	for i := range b {
		b[i] = byte(rand.Intn(maxQual-minQual) + minQual)
	}
	quality = string(b)

	record := strings.Join([]string{header, sequence, repeatHeader, quality}, "\n")

	return []byte(record)
}

// TestGenMultiFastq generates a random valid multiple-record FASTQ
// It generates n 80-length lines using the letters from alphabet.Interface
func TestGenMultiFastq(seed int64, n, m uint, a alphabet.Interface) []byte {
	rand.Seed(seed)
	nseqs := rand.Intn(int(m))
	b := make([]byte, 0)
	for i := 0; i < nseqs; i++ {
		b = append(b, TestGenFastq(seed, n, a)...)
	}
	return b
}
