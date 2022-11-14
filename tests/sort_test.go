package tests

import (
	"testing"

	"github.com/marcopeocchi/fazzoletti/slices"
)

func TestIncludesString(t *testing.T) {
	slice := []int{22, 3, 15, 10}
	match := slices.Includes(slice, 3)
	if !match {
		t.FailNow()
	}
}
