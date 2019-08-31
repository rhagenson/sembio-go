package base_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/io/fasta"
	"github.com/bio-ext/bio-go/bio/io/fasta/base"
	"github.com/bio-ext/bio-go/bio/test"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestDnaIupac(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadDnaIupac removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					hashmap.NewDnaIupac(),
				)
				f, err := base.ReadDnaIupac(ioutil.NopCloser(bytes.NewReader(r)))
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

func TestMultiDnaIupac(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadMultiDnaIupac removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenMultiFasta(
					test.Seed,
					n,
					10,
					hashmap.NewDnaIupac(),
				)
				fs, err := base.ReadMultiDnaIupac(ioutil.NopCloser(bytes.NewReader(r)))
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

func ExampleDnaIupac() {
	x, err := os.Open("../testdata/dna_iupac.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadDnaIupac(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%s\n", f.Header(), f.Sequence())
	// Output:
	// >Generated DNA IUPAC #1
	// YHKWMMTKTASCWGWCGCRNHGNDHM-RTNCYTGWCDMDBWDVVAYTCAHATYSMKAHMCABASMVRMMKSSVM-CYTYVTYBRVCWKBGWAMWVNHATCWMCYMGS--WBAATAHVKWGRKMRTBRVHDDYTBDCRKAHSHRYBTR-SSBAYTKTCMBSSHBYCNGHKNTNWATTSABMTYYDBBMKVBYGHMYSRCVK
}

func ExampleDnaIupac_Header() {
	x, err := os.Open("../testdata/dna_iupac.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadDnaIupac(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Header())
	// Output:
	// >Generated DNA IUPAC #1
}

func ExampleDnaIupac_Sequence() {
	x, err := os.Open("../testdata/dna_iupac.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadDnaIupac(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Sequence())
	// Output:
	// YHKWMMTKTASCWGWCGCRNHGNDHM-RTNCYTGWCDMDBWDVVAYTCAHATYSMKAHMCABASMVRMMKSSVM-CYTYVTYBRVCWKBGWAMWVNHATCWMCYMGS--WBAATAHVKWGRKMRTBRVHDDYTBDCRKAHSHRYBTR-SSBAYTKTCMBSSHBYCNGHKNTNWATTSABMTYYDBBMKVBYGHMYSRCVK
}
