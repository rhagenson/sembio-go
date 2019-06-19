package fasta_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/rhagenson/bio-go/bio/sequence"
	"github.com/rhagenson/bio-go/bio/sequence/immutable"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/io/fasta"
	"github.com/rhagenson/bio-go/bio/test"
)

func TestReadSingle(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadSingle removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					alphabet.NewRna(),
				)
				f, err := fasta.ReadSingle(ioutil.NopCloser(bytes.NewReader(r)), func(s string) (sequence.Interface, error) {
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
				r := fasta.TestGenMultiFasta(
					test.Seed,
					n,
					10,
					alphabet.NewRna(),
				)
				fs, err := fasta.ReadMulti(ioutil.NopCloser(bytes.NewReader(r)), func(s string) (sequence.Interface, error) {
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
