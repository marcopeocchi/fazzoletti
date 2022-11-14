package tests

import (
	"testing"

	"github.com/marcopeocchi/fazzoletti/slices"
)

func TestMapInt(t *testing.T) {
	fixture := []int{2, 6, 10, 22}

	slice := []int{1, 3, 5, 11}
	doubled := slices.Map(slice, func(x int) int {
		return x * 2
	})
	for i := range fixture {
		if fixture[i] != doubled[i] {
			t.FailNow()
		}
	}
}
