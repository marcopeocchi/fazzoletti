package slices

// Partition takes a generic slice and divides it into n sub arrays.
//
//	s := []int{1, 2, 4, 5, 16}
//	p := slices.Partition(s, 2)
//	// p: {[2, 4], [8, 10], [32]}
func Partition[T any](arr []T, chunkSize int) (temp [][]T) {
	temp = [][]T{}
	for i := 0; i < len(arr); i += chunkSize {
		if i == len(arr)-1 {
			temp = append(temp, arr[i:])
			return
		}
		if i >= len(arr) {
			return
		}
		temp = append(temp, arr[i:i+chunkSize])
	}
	return
}
