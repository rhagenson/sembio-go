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

func TestProtein(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadProtein removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					hashmap.NewProtein(),
				)
				f, err := base.ReadProtein(ioutil.NopCloser(bytes.NewReader(r)))
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

func TestMultiProtein(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadMultiProtein removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenMultiFasta(
					test.Seed,
					n,
					10,
					hashmap.NewProtein(),
				)
				fs, err := base.ReadMultiProtein(ioutil.NopCloser(bytes.NewReader(r)))
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

func ExampleProtein() {
	x, err := os.Open("../testdata/protein.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadProtein(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%s\n", f.Header(), f.Sequence())
	// Output:
	// >Generated Protein #1
	// HEWKEYFVQKELDPTWVQLYCWYCLFWAMCVWRHIITWAFTHPMHHFNAHGQAGKMMIYTVAFFVSTTIWMVHTRGHPAMPFKPHWCNQYSGAIYKYPYPRLYNCSCGHDGWLCQGHRATQFTLNHYTFWIEPDLPMEMAGYNGTHTSARNSTKWYQDMANRPHREIFQQMKQTSIMDTYQKWTYRKNNAIKCSQRMKQI
}

func ExampleProtein_Header() {
	x, err := os.Open("../testdata/protein.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadProtein(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Header())
	// Output:
	// >Generated Protein #1
}

func ExampleProtein_Sequence() {
	x, err := os.Open("../testdata/protein.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadProtein(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Sequence())
	// Output:
	// HEWKEYFVQKELDPTWVQLYCWYCLFWAMCVWRHIITWAFTHPMHHFNAHGQAGKMMIYTVAFFVSTTIWMVHTRGHPAMPFKPHWCNQYSGAIYKYPYPRLYNCSCGHDGWLCQGHRATQFTLNHYTFWIEPDLPMEMAGYNGTHTSARNSTKWYQDMANRPHREIFQQMKQTSIMDTYQKWTYRKNNAIKCSQRMKQI
}
