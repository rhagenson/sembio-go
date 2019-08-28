package fastq

import (
	"io"
	"strings"

	"github.com/bio-ext/bio-go/bio/sequence"
)

// Write n records to a FASTQ using the generator f to validate the sequences
// Only records up to the first error are written, returning the number written along with the error
func Write(w io.Writer, is []Interface, n uint, f sequence.Generator) (uint, error) {
	head, seq := "", ""
	heads := make([]string, len(is))
	seqs := make([]string, len(is))
	err, outLength := error(nil), 0
	for i, s := range is {
		if n != 0 && uint(i) == n {
			break
		}
		head = s.Header()
		seq = s.Sequence()
		if _, err = f(seq); err != nil {
			break
		} else {
			heads[i] = head
			seqs[i] = seq
			outLength = outLength + len(head) + len("\n") + len(seq) + len("\n")
		}
	}

	nRecs := uint(0)
	out := new(strings.Builder)
	out.Grow(outLength)
	for i := range heads {
		if heads[i] == "" || seqs[i] == "" {
			break
		}
		out.WriteString(heads[i] + "\n" + seqs[i] + "\n")
		nRecs++
	}
	w.Write([]byte(out.String()))
	return nRecs, err
}

// WriteSingle writes a single record to a FASTQ file using the generator f to validate the sequence
func WriteSingle(w io.Writer, i Interface, f sequence.Generator) (uint, error) {
	return Write(w, []Interface{i}, 1, f)
}

// WriteMulti writes all records to a FASTQ file using the generator f to validate the sequences
// Only records up to the first error are written, returning the number written along with the error
func WriteMulti(w io.Writer, is []Interface, f sequence.Generator) (uint, error) {
	return Write(w, is, 0, f)
}
