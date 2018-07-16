package bigr

import (
	"bytes"
	"encoding/gob"
	"log"
	"math/rand"
	"reflect"
	"testing"
)

// TestSeed is a chosen value that should be used to seed pseudorandom number
// generators
const TestSeed int64 = 1234

// DeepClone does a deep copy from one src to one dest
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

// TestForMethodNamed is a test helper that wraps a check for method by name
// Ideally I would use interfaces to test this property, but interfaces
// match by function signature, not by interface so:
//    `sequence.Interface != *Dna`
//     even if
//    `var _ sequence.Interface = NewDna("")`
func TestForMethodNamed(s interface{}, m string) func(t *testing.T) {
	return func(t *testing.T) {
		if !reflect.ValueOf(s).MethodByName(m).IsValid() {
			t.Errorf("Missing %s method", m)
		}
	}
}

// TestMethodReturnsType checks that s.m(c...) returns type r
func TestMethodReturnsType(s, r interface{}, m string, c []interface{}) func(t *testing.T) {
	cl := make([]reflect.Value, len(c))
	for i := range cl {
		cl[i] = reflect.ValueOf(c[i])
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

// TestMethodReturnsSelfType checks that calling s.m(c) return type s
func TestMethodReturnsSelfType(s interface{}, m string, c []interface{}) func(t *testing.T) {
	return TestMethodReturnsType(s, s, m, c)
}
