package tests

import (
	"testing"

	"github.com/marcopeocchi/fazzoletti/slices"
)

func TestReduceInt(t *testing.T) {
	slice := []int{2, 6, 10, 22, 29, 5}
	sum := slices.Reduce(slice, func(c, n int) int {
		return c + n
	}, 0)
	if sum != 74 {
		t.FailNow()
	}
}
