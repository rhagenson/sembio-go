package immutable_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/sequence"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
	"github.com/bio-ext/bio-go/bio/test"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedDnaIupac(t *testing.T) {
	s, _ := immutable.NewDnaIupac("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestDnaIupacHasMethods(t *testing.T) {
	s, _ := immutable.NewDnaIupac("")

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
	t.Run("Has Alphabet method", func(t *testing.T) {
		if a := s.Alphabet(); a == nil {
			t.Error("Alphabet method does not exist")
		}
	})
	t.Run("Has LetterCount method", func(t *testing.T) {
		if c := s.LetterCount(); c == nil {
			t.Error("LetterCount method does not exist")
		}
	})
}

func TestDnaIupacCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("DnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("DnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("DnaIupac has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
				onefourth := n / 4
				threefourths := n * 3 / 4
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("DnaIupac has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
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
	properties.Property("DnaIupac has proper letter count",
		prop.ForAll(
			func(n uint) bool {
				w := n / 4
				s := test.RandomWeightedString(
					test.Seed,
					n,
					map[rune]uint{
						'A': w,
						'T': w,
						'G': w,
						'C': w,
					},
				)
				seq, _ := immutable.NewDnaIupac(s)
				cs := seq.LetterCount()
				for l, c := range cs {
					if c != w {
						t.Errorf("Got %d, want %d for %q.\nSeq: %q\nMap: %v",
							c, w, l, seq, cs)
						return false
					}
				}
				return true
			},
			gen.UIntRange(4, sequence.TestableLength).
				SuchThat(func(u uint) bool { return u%4 == 0 }),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacImmutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.With(immutable.PositionAs(n/2, t))
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("WithRange does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.With(immutable.RangeAs(n/4, n*3/4, t))
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Reverse does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.Reverse()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.Complement()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.RevComp()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				want, _ := immutable.NewDnaIupac(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*immutable.DnaIupac).Reverse()
				return want.String() == got.(*immutable.DnaIupac).String()
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
					[]rune(hashmap.NewDnaIupac().String()),
				)
				want, _ := immutable.NewDnaIupac(s)
				rev, _ := want.Complement()
				got, _ := rev.(*immutable.DnaIupac).Complement()
				return want.String() == got.(*immutable.DnaIupac).String()
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
					[]rune(hashmap.NewDnaIupac().String()),
				)
				orig, _ := immutable.NewDnaIupac(s)
				want := orig.String()
				rev, _ := orig.RevComp()
				drev, _ := rev.(*immutable.DnaIupac).RevComp()
				got, _ := drev.Range(0, drev.Length())
				return want == got
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(string(alphabet.TestExcludesSingleLetters([]byte(hashmap.NewDnaIupac().String())))),
				)
				if _, err := immutable.NewDnaIupac(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("DnaIupac creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("DnaIupac should error when using invalid characters")
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
					[]rune(hashmap.NewDnaIupac().String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("DnaIupac should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("DnaIupac Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("NewDnaIupac(s) == NewDnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDnaIupac().String()),
				)
				ret := make(chan *immutable.DnaIupac)
				go func(s string, out chan *immutable.DnaIupac) {
					seq, _ := immutable.NewDnaIupac(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *immutable.DnaIupac) {
					seq, _ := immutable.NewDnaIupac(s)
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
					[]rune(hashmap.NewDnaIupac().String()),
				)
				ret := make(chan *immutable.DnaIupac)
				seq, _ := immutable.NewDnaIupac(s)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
				}(seq, ret)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
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
					[]rune(hashmap.NewDnaIupac().String()),
				)
				ret := make(chan *immutable.DnaIupac)
				seq, _ := immutable.NewDnaIupac(s)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
				}(seq, ret)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
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
					[]rune(hashmap.NewDnaIupac().String()),
				)
				ret := make(chan *immutable.DnaIupac)
				seq, _ := immutable.NewDnaIupac(s)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
				}(seq, ret)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
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

// Building a new DnaIupac from valid letters results in no error
func ExampleNewDnaIupac_errorless() {
	s, err := immutable.NewDnaIupac("RYSWKM" + "BDHV" + "N" + "ATGC" + "-")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// RYSWKMBDHVNATGC-, <nil>
}

// Building a new DnaIupac from invalid letters results in an error
// Note that only the first error is returned, not all errors
// The invalid '%' is caught, but nothing is said of the invalid '&'
func ExampleNewDnaIupac_errored() {
	s, err := immutable.NewDnaIupac("%" + "RYSWKM" + "BDHV" + "N" + "ATGC" + "-" + "&")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// %RYSWKMBDHVNATGC-&, "%" not in alphabet
}

// Reversing a valid DnaIupac results in no error
func ExampleDnaIupac_Reverse_errorless() {
	s, _ := immutable.NewDnaIupac("RYSWKM" + "BDHV" + "N" + "ATGC" + "-")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// -CGTANVHDBMKWSYR, <nil>
}

// Reversing an invalid DnaIupac results in an error
// Note that only the first error is returned, not all errors
// The invalid '&' is caught, but nothing is said of the invalid '%'
func ExampleDnaIupac_Reverse_errored() {
	s, _ := immutable.NewDnaIupac("%" + "RYSWKM" + "BDHV" + "N" + "ATGC" + "-" + "&")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// &-CGTANVHDBMKWSYR%, "&" not in alphabet
}

// Reverse complementing a valid DnaIupac results in no error
func ExampleDnaIupac_RevComp_errorless() {
	s, _ := immutable.NewDnaIupac("RYSWKM" + "BDHV" + "N" + "ATGC" + "-")
	rev, err := s.RevComp()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// -GCATNBDHVKMWSRY, <nil>
}

// Reverse complementing an invalid DnaIupac results in an error
// Note that both invalid letters '%' and '&' became 'X' (which is also an invalid letter)
func ExampleDnaIupac_RevComp_errored() {
	s, err := immutable.NewDnaIupac("%" + "RYSWKM" + "BDHV" + "N" + "ATGC" + "-" + "&")
	rev, err := s.RevComp()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// X-GCATNBDHVKMWSRYX, "X" not in alphabet
}

// Complementing a valid DnaIupac results in no error
func ExampleDnaIupac_Complement_errorless() {
	s, _ := immutable.NewDnaIupac("RYSWKM" + "BDHV" + "N" + "ATGC" + "-")
	rev, err := s.Complement()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// YRSWMKVHDBNTACG-, <nil>
}

// Complementing an invalid DnaIupac results in an error
// Note that both invalid letters '%' and '&' became 'X' (which is also an invalid letter)
func ExampleDnaIupac_Complement_errored() {
	s, err := immutable.NewDnaIupac("%" + "RYSWKM" + "BDHV" + "N" + "ATGC" + "-" + "&")
	rev, err := s.Complement()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// XYRSWMKVHDBNTACG-X, "X" not in alphabet
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to immutable.NewDnaIupac()
func ExampleDnaIupac_Alphabet() {
	s, _ := immutable.NewDnaIupac("RYSWKM" + "BDHV" + "N" + "ATGC" + "-")

	fmt.Println(s.Alphabet())
	// Output:
	// -ABCDGHKMNRSTVWY
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to immutable.NewDnaIupac()
func ExampleDnaIupac_LetterCount() {
	s, _ := immutable.NewDnaIupac("RYSWKM" + "BDHV" + "N" + "ATGC" + "-" + "NNNN")

	fmt.Println(s.LetterCount())
	// Output:
	// map[-:1 A:1 B:1 C:1 D:1 G:1 H:1 K:1 M:1 N:5 R:1 S:1 T:1 V:1 W:1 Y:1]
}
