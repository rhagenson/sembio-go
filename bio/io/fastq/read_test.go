package fastq_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sembio/go/bio/sequence"
	"github.com/sembio/go/bio/sequence/immutable"

	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/io/fastq"
	"github.com/sembio/go/bio/test"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestReadSingle(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadSingle removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fastq.TestGenFastq(
					test.Seed,
					n,
					hashmap.NewRna(),
				)
				f, err := fastq.ReadSingle(ioutil.NopCloser(bytes.NewReader(r)), func(s string) (sequence.Interface, error) {
					return immutable.New(s), nil
				})
				if strings.Count(f.Sequence(), "\n") > 1 {
					t.Errorf("body contains internal newline characters: %v", err)
					return false
				}
				return true
			},
			gen.UIntRange(1, 100),
		),
	)
	properties.TestingRun(t)
}

func TestReadMulti(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadMulti removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fastq.TestGenMultiFastq(
					test.Seed,
					n,
					10,
					hashmap.NewRna(),
				)
				fs, err := fastq.ReadMulti(ioutil.NopCloser(bytes.NewReader(r)), func(s string) (sequence.Interface, error) {
					return immutable.New(s), nil
				})
				for _, f := range fs {
					if strings.Count(f.Sequence(), "\n") > 1 {
						t.Errorf("body contains internal newline characters: %v", err)
						return false

					}
				}
				return true
			},
			gen.UIntRange(1, 100),
		),
	)
	properties.TestingRun(t)
}
