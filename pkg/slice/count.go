package slice

// Count returns the number of times val is in s.
func Count[T comparable](s []T, val T) int {
	return CountFunc(s, func(v T) bool { return v == val })
}

// CountFunc returns the number of times f returns true for an element in s.
func CountFunc[T any](s []T, f func(T) bool) int {
	count := 0
	for _, v := range s {
		if f(v) {
			count++
		}
	}
	return count
}
