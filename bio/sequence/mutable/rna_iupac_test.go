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

func TestInitializedRnaIupac(t *testing.T) {
	s, _ := mutable.NewRnaIupac("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestRnaIupacHasMethods(t *testing.T) {
	s, _ := mutable.NewRnaIupac("")

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

func TestRnaIupacCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("RnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := mutable.NewRnaIupac(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := mutable.NewRnaIupac(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := mutable.NewRnaIupac(s)
				onefourth := n / 4
				threefourths := n * 3 / 4
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := mutable.NewRnaIupac(s)
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

func TestRnaIupacMutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := mutable.NewRnaIupac(s)
				clone, _ := mutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := mutable.NewRnaIupac(s)
				clone, _ := mutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				if s == utils.ReverseByBytes(s) { // Skip palindromes
					return true
				}
				original, _ := mutable.NewRnaIupac(s)
				clone, _ := mutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := mutable.NewRnaIupac(s)
				clone, _ := mutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := mutable.NewRnaIupac(s)
				clone, _ := mutable.NewRnaIupac(s)
				original.RevComp()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				want, _ := mutable.NewRnaIupac(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*mutable.RnaIupac).Reverse()
				return want.String() == got.(*mutable.RnaIupac).String()
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				want, _ := mutable.NewRnaIupac(s)
				rev, _ := want.Complement()
				got, _ := rev.(*mutable.RnaIupac).Complement()
				return want.String() == got.(*mutable.RnaIupac).String()
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				want, _ := mutable.NewRnaIupac(s)
				rev, _ := want.RevComp()
				got, _ := rev.(*mutable.RnaIupac).RevComp()
				return want.String() == got.(*mutable.RnaIupac).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(string(alphabet.TestExcludesSingleLetters([]byte(hashmap.NewRnaIupac().String())))),
				)
				if _, err := mutable.NewRnaIupac(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("RnaIupac creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("RnaIupac should error when using invalid characters")
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := mutable.NewRnaIupac(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("RnaIupac should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("RnaIupac Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("mutable.NewRnaIupac(s) == mutable.NewRnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				ret := make(chan *mutable.RnaIupac)
				go func(s string, out chan *mutable.RnaIupac) {
					seq, _ := mutable.NewRnaIupac(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *mutable.RnaIupac) {
					seq, _ := mutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				ret := make(chan *mutable.RnaIupac)
				seq, _ := mutable.NewRnaIupac(s)
				go func(seq *mutable.RnaIupac, out chan *mutable.RnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.RnaIupac)
				}(seq, ret)
				go func(seq *mutable.RnaIupac, out chan *mutable.RnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.RnaIupac)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				ret := make(chan *mutable.RnaIupac)
				seq, _ := mutable.NewRnaIupac(s)
				go func(seq *mutable.RnaIupac, out chan *mutable.RnaIupac) {
					rev, _ := seq.RevComp()
					out <- rev.(*mutable.RnaIupac)
				}(seq, ret)
				go func(seq *mutable.RnaIupac, out chan *mutable.RnaIupac) {
					rev, _ := seq.RevComp()
					out <- rev.(*mutable.RnaIupac)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				ret := make(chan *mutable.RnaIupac)
				seq, _ := mutable.NewRnaIupac(s)
				go func(seq *mutable.RnaIupac, out chan *mutable.RnaIupac) {
					rev, _ := seq.Complement()
					out <- rev.(*mutable.RnaIupac)
				}(seq, ret)
				go func(seq *mutable.RnaIupac, out chan *mutable.RnaIupac) {
					rev, _ := seq.Complement()
					out <- rev.(*mutable.RnaIupac)
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

// Building a new RnaIupac from valid letters results in no error
func ExampleNewRnaIupac_errorless() {
	s, err := mutable.NewRnaIupac("RYSWKM" + "BDHV" + "N" + "AUGC" + "-")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// RYSWKMBDHVNAUGC-, <nil>
}

// Building a new RnaIupac from invalid letters results in an error
// Note that only the first error is returned, not all errors
// The invalid '%' is caught, but nothing is said of the invalid '&'
func ExampleNewRnaIupac_errored() {
	s, err := mutable.NewRnaIupac("%" + "RYSWKM" + "BDHV" + "N" + "AUGC" + "-" + "&")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// %RYSWKMBDHVNAUGC-&, "%" not in alphabet
}

// Reversing a valid RnaIupac results in no error
func ExampleRnaIupac_Reverse_errorless() {
	s, _ := mutable.NewRnaIupac("RYSWKM" + "BDHV" + "N" + "AUGC" + "-")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// -CGUANVHDBMKWSYR, <nil>
}

// Reversing an invalid RnaIupac results in an error
// Note that only the first error is returned, not all errors
// The invalid '&' is caught, but nothing is said of the invalid '%'
func ExampleRnaIupac_Reverse_errored() {
	s, _ := mutable.NewRnaIupac("%" + "RYSWKM" + "BDHV" + "N" + "AUGC" + "-" + "&")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// &-CGUANVHDBMKWSYR%, "&" not in alphabet
}

// Reverse complementing a valid RnaIupac results in no error
func ExampleRnaIupac_RevComp_errorless() {
	s, _ := mutable.NewRnaIupac("RYSWKM" + "BDHV" + "N" + "AUGC" + "-")
	rev, err := s.RevComp()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// -GCAUNBDHVKMWSRY, <nil>
}

// Reverse complementing an invalid RnaIupac results in an error
// Note that both invalid letters '%' and '&' became 'X' (which is also an invalid letter)
func ExampleRnaIupac_RevComp_errored() {
	s, err := mutable.NewRnaIupac("%" + "RYSWKM" + "BDHV" + "N" + "AUGC" + "-" + "&")
	rev, err := s.RevComp()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// X-GCAUNBDHVKMWSRYX, "X" not in alphabet
}

// Complementing a valid RnaIupac results in no error
func ExampleRnaIupac_Complement_errorless() {
	s, _ := mutable.NewRnaIupac("RYSWKM" + "BDHV" + "N" + "AUGC" + "-")
	rev, err := s.Complement()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// YRSWMKVHDBNUACG-, <nil>
}

// Complementing an invalid RnaIupac results in an error
// Note that both invalid letters '%' and '&' became 'X' (which is also an invalid letter)
func ExampleRnaIupac_Complement_errored() {
	s, err := mutable.NewRnaIupac("%" + "RYSWKM" + "BDHV" + "N" + "AUGC" + "-" + "&")
	rev, err := s.Complement()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// XYRSWMKVHDBNUACG-X, "X" not in alphabet
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewRnaIupac()
func ExampleRnaIupac_Alphabet() {
	s, _ := mutable.NewRnaIupac("RYSWKM" + "BDHV" + "N" + "AUGC" + "-")

	fmt.Println(s.Alphabet())
	// Output:
	// -ABCDGHKMNRSUVWY
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewRnaIupac()
func ExampleRnaIupac_LetterCount() {
	s, _ := mutable.NewRnaIupac("RYSWKM" + "BDHV" + "N" + "AUGC" + "-" + "NNNN")

	fmt.Println(s.LetterCount())
	// Output:
	// map[-:1 A:1 B:1 C:1 D:1 G:1 H:1 K:1 M:1 N:5 R:1 S:1 U:1 V:1 W:1 Y:1]
}
