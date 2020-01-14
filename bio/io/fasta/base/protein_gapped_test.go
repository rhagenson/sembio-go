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

func TestProteinGapped(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadProteinGapped removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					hashmap.NewProteinGapped(),
				)
				f, err := base.ReadProteinGapped(ioutil.NopCloser(bytes.NewReader(r)))
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

func TestMultiProteinGapped(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadMultiProteinGapped removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenMultiFasta(
					test.Seed,
					n,
					10,
					hashmap.NewProteinGapped(),
				)
				fs, err := base.ReadMultiProteinGapped(ioutil.NopCloser(bytes.NewReader(r)))
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

func ExampleProteinGapped() {
	x, err := os.Open("../testdata/protein_gapped.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadProteinGapped(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%s\n", f.Header(), f.Sequence())
	// Output:
	// >Generated Protein Gapped #1
	// HEWKEYFVQKELDPT---LYCWYCLFWAMCVWRHIITWAF-HPMHHFNAHGQAGKMMIYTVAFFVSTTIWMVHTRGH-AMPFKPHWCNQYSGAIYKYPYP-LYNCSCGHDGWLCQGHRATQFTLNHYTFWIEPDLPM-MAGYNGTHTSARNSTKWYQDMA-RPHREIFQQMKQTSIMDTYQKWTYRKNNAIKCSQRM-QI
}

func ExampleProteinGapped_Header() {
	x, err := os.Open("../testdata/protein_gapped.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadProteinGapped(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Header())
	// Output:
	// >Generated Protein Gapped #1
}

func ExampleProteinGapped_Sequence() {
	x, err := os.Open("../testdata/protein_gapped.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadProteinGapped(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Sequence())
	// Output:
	// HEWKEYFVQKELDPT---LYCWYCLFWAMCVWRHIITWAF-HPMHHFNAHGQAGKMMIYTVAFFVSTTIWMVHTRGH-AMPFKPHWCNQYSGAIYKYPYP-LYNCSCGHDGWLCQGHRATQFTLNHYTFWIEPDLPM-MAGYNGTHTSARNSTKWYQDMA-RPHREIFQQMKQTSIMDTYQKWTYRKNNAIKCSQRM-QI
}
