package slices

// Array.prototype.filter() inspired function.
// Filter takes a slices of type T and a certain value and
// filters it down to the elements that pass the test function.
//
//	a := []string{"ciao", "come", "va", "grande"}
//	filtered := slices.Filter(a, func(e string) bool {
//		return strings.HasPrefix(e, "c")
//	})
//	//filtered: {"ciao", "come"}
func Filter[T any](slice []T, f func(e T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
