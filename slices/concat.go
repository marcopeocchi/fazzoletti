package slices

// Array.prototype.concat inspired function
// Concat takes a slice of slices and concat them without
// using append() function.
func Concat[T any](s [][]T) []T {
	var (
		i        int
		totalLen int
	)

	for _, s := range s {
		totalLen += len(s)
	}

	result := make([]T, totalLen)
	for _, s := range s {
		i += copy(result[i:], s)
	}

	return result
}
