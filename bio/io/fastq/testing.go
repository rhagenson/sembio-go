package fastq

import (
	"math/rand"
	"strings"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/data/quality"
)

// TestGenFastq generates a random valid FASTQ
func TestGenFastq(seed int64, n uint, a alphabet.Interface) []byte {
	rand.Seed(seed)
	valid := a.String()
	linelen := func() uint {
		if n < 80 {
			return n / 2
		}
		return 80
	}()

	header := string(FastqHeaderPrefix) + "EAS139:136:FC706VJ:2:2104:15343:197393 1:Y:18:ATCACG"
	sequence := ""
	repeatHeader := string(FastqPreQualityHeaderPrefix) + header[1:]
	qualityLine := ""

	b := make([]byte, linelen)
	for i := range b {
		b[i] = valid[rand.Intn(len(valid))]
	}
	sequence = string(b)

	minQual := int(quality.MinPhred33)
	maxQual := int(quality.MaxIlluminaPhred33)
	for i := range b {
		b[i] = byte(rand.Intn(maxQual-minQual) + minQual)
	}
	qualityLine = string(b)

	record := strings.Join([]string{header, sequence, repeatHeader, qualityLine}, "\n")

	return []byte(record)
}

// TestGenMultiFastq generates a random valid multiple-record FASTQ
func TestGenMultiFastq(seed int64, n, m uint, a alphabet.Interface) []byte {
	rand.Seed(seed)
	nseqs := rand.Intn(int(m))
	b := make([]byte, 0)
	for i := 0; i < nseqs; i++ {
		b = append(b, TestGenFastq(seed, n, a)...)
	}
	return b
}
