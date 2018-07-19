package sequence

import (
	"reflect"
	"testing"
)

// TestableLength is a realistic, testable length for
// generating random sequences in tests
const TestableLength uint = 1000

// TestAlphabetIs is a test helper that wraps a check for expected alphabet
func TestAlphabetIs(s interface{}, exp interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		if !reflect.DeepEqual(exp, s) {
			t.Errorf("Want: %q, Got: %q", exp, s)
		}
	}
}

// TestLengthIs is a test helper that wraps a check for known length
func TestLengthIs(s Interface, exp uint) func(t *testing.T) {
	return func(t *testing.T) {
		if !reflect.DeepEqual(exp, s.Length()) {
			t.Errorf("Want: %q, Got: %q", exp, s.Length())
		}
	}
}

// TestPositionIs is a test helper that wraps a check for known position
func TestPositionIs(s Interface, p uint, exp string) func(t *testing.T) {
	return func(t *testing.T) {
		got, _ := s.Position(p)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Want: %q, Got: %q", exp, got)
		}
	}
}

// TestRangeIs is a test helper that wraps a check for known range
func TestRangeIs(s Interface, st, sp uint, exp string) func(t *testing.T) {
	return func(t *testing.T) {
		got, _ := s.Range(st, sp)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Want: %q, Got: %q", exp, got)
		}
	}
}
