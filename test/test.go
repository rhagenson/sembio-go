package test

import (
	"bytes"
	"encoding/gob"
	"log"
	"math/rand"
	"reflect"
	"testing"
)

// Seed is a chosen value that should be used to
// seed pseudorandom number generators
const Seed int64 = 1234

// DeepClone does a deep copy from one src to one dest
// Note: DeepClone copies only the public parts of a struct
func DeepClone(src, dest interface{}) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	dec := gob.NewDecoder(&buff)
	err := enc.Encode(src)
	if err != nil {
		log.Fatal("encode error: ", err)
	}
	err = dec.Decode(&dest)
	if err != nil {
		log.Fatal("decode error: ", err)
	}
}

// RandomStringFromRunes generates a random string of length n from
// the supplied valid runes
func RandomStringFromRunes(seed int64, n uint, valid []rune) string {
	rand.Seed(seed)
	b := make([]rune, n)
	for i := range b {
		b[i] = valid[rand.Intn(len(valid))]
	}
	return string(b)
}

// RandomWeightedString generates a random weighted string of length n.
// If n != sum(weights), an approximation is used
func RandomWeightedString(seed int64, n uint, weights map[rune]uint) string {
	rand.Seed(seed)

	// Allocate memory
	s := make([]rune, n)

	// Fill with valid letters dependent on weights
	idx := 0
outer:
	for r, w := range weights {
		for i := idx; i < idx+int(w); i++ {
			if i < int(n) {
				s[i] = r
			} else {
				break outer
			}
		}
		idx = idx + int(w)
	}

	// Shuffle array
	for i := range s {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func gcd_poly(ns ...int) int {
	n := len(ns)
	if n == 1 {
		return ns[0]
	}
	if n == 2 {
		return gcd(ns[0], ns[1])
	}
	h := n / 2
	return gcd(gcd_poly(ns[:h]...), gcd_poly(ns[h:]...))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// ForMethodNamed is a test helper that wraps a check for method by name
func ForMethodNamed(s interface{}, m string) func(t *testing.T) {
	return func(t *testing.T) {
		if !reflect.ValueOf(s).MethodByName(m).IsValid() {
			t.Errorf("Missing %s method", m)
		}
	}
}

// MethodReturnsType checks that s.m(args...) returns type r
func MethodReturnsType(s, r interface{}, m string, args []interface{}) func(t *testing.T) {
	cl := make([]reflect.Value, len(args))
	for i := range cl {
		cl[i] = reflect.ValueOf(args[i])
	}
	return func(t *testing.T) {
		want := reflect.TypeOf(r).String()
		for _, v := range reflect.ValueOf(s).MethodByName(m).Call(cl) {
			got := v.Type().String()
			if want != got {
				t.Errorf("Got %q; Want %q",
					got,
					want,
				)
			}
		}
	}
}

// MethodReturnsSelfType checks that calling s.m(args...) return type s
func MethodReturnsSelfType(s interface{}, m string, args []interface{}) func(t *testing.T) {
	return MethodReturnsType(s, s, m, args)
}
