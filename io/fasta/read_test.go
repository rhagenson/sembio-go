package fasta_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/immutable"

	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/io/fasta"
	"bitbucket.org/rhagenson/bio/test"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestRead(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Read removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					alphabet.Rna,
				)

				ch1, ch2 := fasta.Read(ioutil.NopCloser(bytes.NewReader(r)), func(s string) (sequence.Interface, error) {
					seq := immutable.New(s)
					return seq, seq.Validate()
				}, 1)
				f := <-ch1
				err := <-ch2
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

func TestReadSingle(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadSingle removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					alphabet.Rna,
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
