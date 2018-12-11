package fasta

import (
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bio"
	"bitbucket.org/rhagenson/bio/alphabet"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestRnaIupac(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadRnaIupa removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := TestGenFasta(
					bio.TestSeed,
					n,
					alphabet.RnaIupac,
				)
				f, err := ReadRnaIupac(r)
				if strings.Count(f.Body(), "\n") > 1 {
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
