package slice

// MinFunc returns the minimum value of a slice.
// The function min is used to compare the values.
// The function min should return true if the first value is smaller than the second value.
func MinFunc[T any](s []T, min func(T, T) bool) T {
	minValue := s[0]
	for _, v := range s {
		if min(v, minValue) {
			minValue = v
		}
	}
	return minValue
}
