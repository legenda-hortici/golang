package main

import (
	"math/rand"
	"testing"
)

func TestCreate(t *testing.T) {
	rand.Seed(42)

	n := 10
	res := create(n)

	seen := make(map[int]bool)
	for _, v := range res {
		if seen[v] {
			t.Fatalf("duplicate found: %d in %v", v, res)
		}
		seen[v] = true
	}

	if len(res) > n {
		t.Fatalf("length too big: got %d, want <= %d", len(res), n)
	}

	for _, v := range res {
		if v < 0 || v >= 20 {
			t.Fatalf("value out of range: %d", v)
		}
	}
}
