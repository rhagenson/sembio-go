package mutable_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sembio/go/bio/alphabet"
	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/sequence"
	"github.com/sembio/go/bio/sequence/mutable"
	"github.com/sembio/go/bio/test"
	"github.com/sembio/go/bio/utils"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedRna(t *testing.T) {
	s, _ := mutable.NewRna("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestRnaHasMethods(t *testing.T) {
	s, _ := mutable.NewRna("")

	t.Run("Has Reverse method", func(t *testing.T) {
		if _, err := s.Reverse(); err != nil {
			t.Error("Reverse method does not exist")
		}
	})
	t.Run("Has Complement method", func(t *testing.T) {
		if _, err := s.Complement(); err != nil {
			t.Error("Complement method does not exist")
		}
	})
	t.Run("Has RevComp method", func(t *testing.T) {
		if _, err := s.RevComp(); err != nil {
			t.Error("RevComp method does not exist")
		}
	})
}

func TestRnaCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Rna is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				seq, _ := mutable.NewRna(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Rna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				seq, _ := mutable.NewRna(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Rna has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				seq, _ := mutable.NewRna(s)
				onefourth := n / 4
				threefourths := n * 3 / 4
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Rna has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				seq, _ := mutable.NewRna(s)
				onefourth := n / 4
				threefourth := n * (3 / 4)
				gotoneforth, _ := seq.Position(onefourth)
				wantoneforth := string(s[onefourth])
				gotthreeforth, _ := seq.Position(threefourth)
				wantthreeforth := string(s[threefourth])
				return gotoneforth == wantoneforth && gotthreeforth == wantthreeforth
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaMutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				original, _ := mutable.NewRna(s)
				clone, _ := mutable.NewRna(s)
				original.With(mutable.PositionAs(n/2, t))
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("WithRange mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				original, _ := mutable.NewRna(s)
				clone, _ := mutable.NewRna(s)
				original.With(mutable.RangeAs(n/4, n*3/4, t))
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Reverse mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				if s == utils.ReverseByBytes(s) { // Skip palindromes
					return true
				}
				original, _ := mutable.NewRna(s)
				clone, _ := mutable.NewRna(s)
				original.Reverse()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				original, _ := mutable.NewRna(s)
				clone, _ := mutable.NewRna(s)
				original.Complement()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				original, _ := mutable.NewRna(s)
				clone, _ := mutable.NewRna(s)
				original.RevComp()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				want, _ := mutable.NewRna(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*mutable.Rna).Reverse()
				return want.String() == got.(*mutable.Rna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement().Complement() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				want, _ := mutable.NewRna(s)
				rev, _ := want.Complement()
				got, _ := rev.(*mutable.Rna).Complement()
				return want.String() == got.(*mutable.Rna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp().RevComp() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				want, _ := mutable.NewRna(s)
				rev, _ := want.RevComp()
				got, _ := rev.(*mutable.Rna).RevComp()
				return want.String() == got.(*mutable.Rna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(string(alphabet.TestExcludesSingleLetters([]byte(hashmap.NewRna().String())))),
				)
				if _, err := mutable.NewRna(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("Rna creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("Rna should error when using invalid characters")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("start > stop errors",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				seq, _ := mutable.NewRna(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("Rna should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("Rna Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("mutable.NewRna(s) == mutable.NewRna(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				ret := make(chan *mutable.Rna)
				go func(s string, out chan *mutable.Rna) {
					seq, _ := mutable.NewRna(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *mutable.Rna) {
					seq, _ := mutable.NewRna(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.Reverse() == seq.Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				ret := make(chan *mutable.Rna)
				seq, _ := mutable.NewRna(s)
				go func(seq *mutable.Rna, out chan *mutable.Rna) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.Rna)
				}(seq, ret)
				go func(seq *mutable.Rna, out chan *mutable.Rna) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.Rna)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.RevComp() == seq.RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				ret := make(chan *mutable.Rna)
				seq, _ := mutable.NewRna(s)
				go func(seq *mutable.Rna, out chan *mutable.Rna) {
					rev, _ := seq.RevComp()
					out <- rev.(*mutable.Rna)
				}(seq, ret)
				go func(seq *mutable.Rna, out chan *mutable.Rna) {
					rev, _ := seq.RevComp()
					out <- rev.(*mutable.Rna)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.Complement() == seq.Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRna().String()),
				)
				ret := make(chan *mutable.Rna)
				seq, _ := mutable.NewRna(s)
				go func(seq *mutable.Rna, out chan *mutable.Rna) {
					rev, _ := seq.Complement()
					out <- rev.(*mutable.Rna)
				}(seq, ret)
				go func(seq *mutable.Rna, out chan *mutable.Rna) {
					rev, _ := seq.Complement()
					out <- rev.(*mutable.Rna)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

// Building a new Rna from valid letters results in no error
func ExampleNewRna_errorless() {
	s, err := mutable.NewRna("AUGC")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// AUGC, <nil>
}

// Building a new Rna from invalid letters results in an error
// Note that only the first error is returned, not all errors
// The invalid '%' is caught, but nothing is said of the invalid '&'
func ExampleNewRna_errored() {
	s, err := mutable.NewRna("%" + "AUGC" + "&")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// %AUGC&, "%" not in alphabet
}

// Reversing a valid Rna results in no error
func ExampleRna_Reverse_errorless() {
	s, _ := mutable.NewRna("AUGC")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// CGUA, <nil>
}

// Reversing an invalid Rna results in an error
// Note that only the first error is returned, not all errors
// The invalid '&' is caught, but nothing is said of the invalid '%'
func ExampleRna_Reverse_errored() {
	s, _ := mutable.NewRna("%" + "AUGC" + "&")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// &CGUA%, "&" not in alphabet
}

// Reverse complementing a valid Rna results in no error
func ExampleRna_RevComp_errorless() {
	s, _ := mutable.NewRna("AUGC")
	rev, err := s.RevComp()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// GCAU, <nil>
}

// Reverse complementing an invalid Rna results in an error
// Note that both invalid letters '%' and '&' became 'X' (which is also an invalid letter)
func ExampleRna_RevComp_errored() {
	s, err := mutable.NewRna("%" + "AUGC" + "&")
	rev, err := s.RevComp()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// XGCAUX, "X" not in alphabet
}

// Complementing a valid Rna results in no error
func ExampleRna_Complement_errorless() {
	s, _ := mutable.NewRna("AUGC")
	rev, err := s.Complement()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// UACG, <nil>
}

// Complementing an invalid Rna results in an error
// Note that both invalid letters '%' and '&' became 'X' (which is also an invalid letter)
func ExampleRna_Complement_errored() {
	s, err := mutable.NewRna("%" + "AUGC" + "&")
	rev, err := s.Complement()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// XUACGX, "X" not in alphabet
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewRna()
func ExampleRna_Alphabet() {
	s, _ := mutable.NewRna("AUGC")

	fmt.Println(s.Alphabet())
	// Output:
	// ACGU
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewRna()
func ExampleRna_LetterCount() {
	s, _ := mutable.NewRna("AUGC" + "AAAA")

	fmt.Println(s.LetterCount())
	// Output:
	// map[A:5 C:1 G:1 U:1]
}
