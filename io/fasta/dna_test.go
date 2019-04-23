package fasta_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/io/fasta"
	"bitbucket.org/rhagenson/bio/test"
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
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					alphabet.Dna,
				)
				f, err := fasta.ReadDna(ioutil.NopCloser(bytes.NewReader(r)))
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
				r := fasta.TestGenMultiFasta(
					test.Seed,
					n,
					alphabet.Dna,
				)
				fs, err := fasta.ReadDna(ioutil.NopCloser(bytes.NewReader(r)))
				// for _, f := range fs {
				switch {
				case strings.Count(fs.Sequence(), "\n") > 1:
					t.Errorf("body contains internal newline characters: %v", err)
					return false
				case err != nil:
					t.Errorf("error in parsing input: %v", err)
					return false
				default:
					return true
				}
				// }
				// return true
			},
			gen.UIntRange(1, 100),
		),
	)
	properties.TestingRun(t)
}
