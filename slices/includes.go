package slices

// Array.prototype.includes() inspired function.
// Includes takes a slices of type T and a certain value and
// determines whether the slice include the value among its entries.
//
//	a := []string{"ciao", "come", "va", "grande"}
//	c1 := slices.Includes(a, "va") // true
//	c2 := slices.Includes(a, "vaa") // false
func Includes[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
