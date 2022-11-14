package slices

// Array.prototype.map() inspired function.
// Map takes a slices of type T and a certain value and
// transform each value with a mapper function.
// The result will be a slice of the same type.
//
//	a := []int{1, 2, 4, 5, 16}
//	mapped := slices.Map(a, func(e int) int {
//		return e * 2
//	})
//	//mapped: {2, 4, 8, 10, 32}
func Map[T any, M any](a []T, f func(x T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}
