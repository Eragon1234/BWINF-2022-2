package slice

// SumFunc sums the slice with the given sum function.
// The sum function takes the current sum and the current element and returns the new sum.
func SumFunc[T any, M any](slice []T, sum func(M, T) M) M {
	var result M
	for _, v := range slice {
		result = sum(result, v)
	}
	return result
}
