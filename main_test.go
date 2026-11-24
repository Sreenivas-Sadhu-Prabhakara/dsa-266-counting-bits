package main

import (
	"math/bits"
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	cases := []struct {
		n    int
		want []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
		{0, []int{0}},
	}
	for _, c := range cases {
		if got := CountBits(c.n); !reflect.DeepEqual(got, c.want) {
			t.Errorf("CountBits(%d) = %v, want %v", c.n, got, c.want)
		}
	}
	// Property: agree with the stdlib popcount everywhere up to 4096.
	for i, got := range CountBits(4096) {
		if got != bits.OnesCount(uint(i)) {
			t.Fatalf("mismatch at %d: %d", i, got)
		}
	}
}
