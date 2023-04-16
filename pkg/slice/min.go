package slice

// MinFunc returns the minimum value of a slice.
// The function min is used to compare the values.
// The function min should return true if the first value is smaller than the second value.
// If the slice is empty, the zero value of the type is returned.
func MinFunc[T any](s []T, min func(T, T) bool) T {
	if len(s) == 0 {
		return *new(T)
	}
	minValue := s[0]
	for _, v := range s {
		if min(v, minValue) {
			minValue = v
		}
	}
	return minValue
}
