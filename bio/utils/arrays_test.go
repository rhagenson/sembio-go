package utils_test

import (
	"fmt"
	"testing"

	"github.com/bio-ext/bio-go/bio/utils"
)

func TestBytesToStrings(t *testing.T) {
	tt := []struct {
		in  []byte
		out []string
	}{
		{[]byte("ABC"), []string{"A", "B", "C"}},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%#v", tc.in), func(t *testing.T) {
			got := utils.BytesToStrings(tc.in)
			for i := range tc.out {
				if got[i] != tc.out[i] {
					t.Errorf("got %q, want %q", got, tc.out)
				}
			}
		})
	}
}

func TestStingsToBytes(t *testing.T) {
	tt := []struct {
		in  []string
		out []byte
	}{
		{[]string{"A", "B", "C"}, []byte("ABC")},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%#v", tc.in), func(t *testing.T) {
			got := utils.StringsToBytes(tc.in)
			for i := range tc.out {
				if got[i] != tc.out[i] {
					t.Errorf("got %q, want %q", got, tc.out)
				}
			}
		})
	}
}
