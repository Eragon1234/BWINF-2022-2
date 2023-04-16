package slice

// MakeFunc creates a slice of length len and fills it with the values returned by f.
// f is called len times with the index as argument.
func MakeFunc[T any](len int, f func(i int) T) []T {
	s := make([]T, len)
	for i := 0; i < len; i++ {
		s[i] = f(i)
	}
	return s
}
