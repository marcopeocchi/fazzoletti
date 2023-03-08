package tests

import (
	"testing"

	"github.com/marcopeocchi/fazzoletti/slices"
)

func TestPartitionInt(t *testing.T) {
	s := []int{2, 6, 10, 22, 29, 5, 8}
	p := slices.Partition(s, 2)
	if len(p) != 4 {
		t.FailNow()
	}
	for _, pp := range p {
		if len(pp) > 2 || len(pp) == 0 {
			t.FailNow()
		}
	}
}

func BenchmarkPartitionInt(b *testing.B) {
	s := make([]int, b.N)
	for i := range s {
		s[i] = i
	}
	_ = slices.Partition(s, 3)
}
