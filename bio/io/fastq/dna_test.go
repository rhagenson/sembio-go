package fastq_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/io/fastq"
	"github.com/sembio/go/bio/test"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestDna(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadDna removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fastq.TestGenFastq(
					test.Seed,
					n,
					hashmap.NewDna(),
				)
				f, err := fastq.ReadDna(ioutil.NopCloser(bytes.NewReader(r)))
				switch {
				case strings.Count(f.Sequence(), "\n") > 1:
					t.Errorf("body contains internal newline characters: %v", err)
					return false
				case err != nil:
					t.Errorf("error in parsing input: %v", err)
					return false
				default:
					return true
				}
			},
			gen.UIntRange(1, 100),
		),
	)
	properties.TestingRun(t)
}

func TestMultiDna(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadMultiDna removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fastq.TestGenMultiFastq(
					test.Seed,
					n,
					10,
					hashmap.NewDna(),
				)
				fs, err := fastq.ReadMultiDna(ioutil.NopCloser(bytes.NewReader(r)))
				for _, f := range fs {
					switch {
					case strings.Count(f.Sequence(), "\n") > 1:
						t.Errorf("body contains internal newline characters: %v", err)
						return false
					case err != nil:
						t.Errorf("error in parsing input: %v", err)
						return false
					default:
						return true
					}
				}
				return true
			},
			gen.UIntRange(100, 1000),
		),
	)
	properties.TestingRun(t)
}
