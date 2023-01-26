package utils

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
