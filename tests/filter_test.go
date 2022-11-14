package tests

import (
	"strings"
	"testing"

	"github.com/marcopeocchi/fazzoletti/slices"
)

func TestFilterString(t *testing.T) {
	slice := []string{"this", "is", "a", "test", "on", "a", "strings", "slice"}
	filtered := slices.Filter(slice, func(e string) bool {
		return strings.HasPrefix(e, "t")
	})
	if len(filtered) != 2 {
		t.Fatal()
	}
}

func TestFilterStruct(t *testing.T) {
	type User struct {
		username string
		role     string
	}

	slice := []User{
		{username: "admin", role: "ADMIN"},
		{username: "user1", role: "USER"},
		{username: "user2", role: "USER"},
	}
	filtered := slices.Filter(slice, func(e User) bool {
		return e.role != "ADMIN"
	})
	if !(len(filtered) == 2) {
		t.Fatal()
	}
}
