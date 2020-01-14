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

func TestRna(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadRna removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenFasta(
					test.Seed,
					n,
					hashmap.NewRna(),
				)
				f, err := base.ReadRna(ioutil.NopCloser(bytes.NewReader(r)))
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

func TestMultiRna(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ReadMultiRna removes newline characters in body",
		prop.ForAll(
			func(n uint) bool {
				r := fasta.TestGenMultiFasta(
					test.Seed,
					n,
					10,
					hashmap.NewRna(),
				)
				fs, err := base.ReadMultiRna(ioutil.NopCloser(bytes.NewReader(r)))
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

func ExampleRna() {
	x, err := os.Open("../testdata/rna.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadRna(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%s\n", f.Header(), f.Sequence())
	// Output:
	// >Generated RNA #1
	// UGAUGCAUGAUAACUACAUGCCUAUAGUUAGUGAAGGAAGGCUGUUCCACAUUGACCGUGCUGCGUACAGAUUCACUGGGUUGAGCAACCCAACGAGGUAGUGUAUGUUGGUUAGUCUAGGAACCCGGUCUCGUGUCGAUGUUUGGGGGGUCGCCGUAAGUAGAAAAUUUCGGUCGAGAUAUCCUUCCAGCUUUUAUCCG
}

func ExampleRna_Header() {
	x, err := os.Open("../testdata/rna.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadRna(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Header())
	// Output:
	// >Generated RNA #1
}

func ExampleRna_Sequence() {
	x, err := os.Open("../testdata/rna.fasta")
	if err != nil {
		panic(err)
	}
	f, err := base.ReadRna(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", f.Sequence())
	// Output:
	// UGAUGCAUGAUAACUACAUGCCUAUAGUUAGUGAAGGAAGGCUGUUCCACAUUGACCGUGCUGCGUACAGAUUCACUGGGUUGAGCAACCCAACGAGGUAGUGUAUGUUGGUUAGUCUAGGAACCCGGUCUCGUGUCGAUGUUUGGGGGGUCGCCGUAAGUAGAAAAUUUCGGUCGAGAUAUCCUUCCAGCUUUUAUCCG
}
