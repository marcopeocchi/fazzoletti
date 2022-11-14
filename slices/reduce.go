package slices

// Array.prototype.reduce() inspired function.
// Reduces executes a "reducer" function on each element of the slice
// passing in the result of the previous execution.
// The final result is a single value.
//
//	a := []int{3, 5, 7, 9}
//	sum := slices.Map(a, func(c int, n int) int {
//		return c + n
//	})
//	//sum: 24
func Reduce[T, U any](s []T, f func(c U, n T) U, initValue U) U {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}
