package slice

// Pop removes the last element from the slice and returns it.
// If the slice is empty, the zero value of the slice's element type is returned.
func Pop[T any](s []T) (T, []T) {
	i := len(s) - 1
	if i < 0 {
		return *new(T), s
	}
	return s[i], s[:i]
}
