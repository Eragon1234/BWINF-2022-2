package slice

func MakeFunc[T any](len int, f func(i int) T) []T {
	s := make([]T, len)
	for i := 0; i < len; i++ {
		s[i] = f(i)
	}
	return s
}
