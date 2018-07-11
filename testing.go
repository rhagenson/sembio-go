package bigr

import (
	"bytes"
	"encoding/gob"
	"log"
	"math/rand"
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
