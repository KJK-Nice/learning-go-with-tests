package slices

// Sum returns the sum of a slice of integers.
func Sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
