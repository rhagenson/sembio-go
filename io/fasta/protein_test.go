package fasta_test

import (
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bio"
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/io/fasta"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestProtein(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadProtein removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					bio.TestSeed,
					n,
					alphabet.Protein,
				)
				f, err := fasta.ReadProtein(r)
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
