package utils

func Pop[T any](s []T) (T, []T) {
	i := len(s) - 1
	if i < 0 {
		return *new(T), s
	}
	return s[i], s[:i]
}
