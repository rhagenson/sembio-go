package base_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/io/fasta"
	"github.com/sembio/go/bio/io/fasta/base"

	"github.com/sembio/go/bio/test"
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
					hashmap.NewRnaIupac(),
				)
				f, err := base.ReadRnaIupac(ioutil.NopCloser(bytes.NewReader(r)))
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

func TestMultiRnaIupac(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadMultiRnaIupac removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenMultiFasta(
					test.Seed,
					n,
					10,
					hashmap.NewRnaIupac(),
				)
				fs, err := base.ReadMultiRnaIupac(ioutil.NopCloser(bytes.NewReader(r)))
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

func ExampleRnaIupac() {
	x, err := os.Open("../testdata/rna_iupac.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadRnaIupac(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%s\n", f.Header(), f.Sequence())
	// Output:
	// >Generated RNA IUPAC #1
	// YHKWMMUKUASCWGWCGCRNHGNDHM-RUNCYUGWCDMDBWDVVAYUCAHAUYSMKAHMCABASMVRMMKSSVM-CYUYVUYBRVCWKBGWAMWVNHAUCWMCYMGS--WBAAUAHVKWGRKMRUBRVHDDYUBDCRKAHSHRYBUR-SSBAYUKUCMBSSHBYCNGHKNUNWAUUSABMUYYDBBMKVBYGHMYSRCVK
}

func ExampleRnaIupac_Header() {
	x, err := os.Open("../testdata/rna_iupac.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadRnaIupac(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Header())
	// Output:
	// >Generated RNA IUPAC #1
}

func ExampleRnaIupac_Sequence() {
	x, err := os.Open("../testdata/rna_iupac.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadRnaIupac(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Sequence())
	// Output:
	// YHKWMMUKUASCWGWCGCRNHGNDHM-RUNCYUGWCDMDBWDVVAYUCAHAUYSMKAHMCABASMVRMMKSSVM-CYUYVUYBRVCWKBGWAMWVNHAUCWMCYMGS--WBAAUAHVKWGRKMRUBRVHDDYUBDCRKAHSHRYBUR-SSBAYUKUCMBSSHBYCNGHKNUNWAUUSABMUYYDBBMKVBYGHMYSRCVK
}
