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

func TestRnaIupac(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadRnaIupac removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					alphabet.RnaIupac,
				)
				f, err := fasta.ReadRnaIupac(ioutil.NopCloser(bytes.NewReader(r)))
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
