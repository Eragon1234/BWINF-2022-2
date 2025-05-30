package slice

// Map returns a new slice with the result of applying f to each element of s.
func Map[T any, U any](s []T, f func(T) U) []U {
	mapped := make([]U, len(s))
	for i, e := range s {
		mapped[i] = f(e)
	}
	return mapped
}
