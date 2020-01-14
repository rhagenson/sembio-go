package fastq_test

import (
	"testing"

	"github.com/sembio/go/bio/io/fastq"
	"github.com/sembio/go/bio/sequence/immutable"
	"github.com/sembio/go/bio/test"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestStruct(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Header is the same as input",
		prop.ForAll(
			func(in string) bool {
				out := fastq.New(in, "", nil).Header()
				if in != out {
					t.Errorf("input, %q, did not match output %q", in, out)
					return false
				}
				return true
			},
			gen.AlphaString(),
		),
	)
	properties.Property("Sequence is the same as input",
		prop.ForAll(
			func(in string) bool {
				seq := immutable.New(in)
				out := fastq.New("", "", seq).Sequence()
				if in != out {
					t.Errorf("input, %q, did not match output %q", in, out)
					return false
				}
				return true
			},
			gen.AlphaString(),
		),
	)
	properties.TestingRun(t)
}
