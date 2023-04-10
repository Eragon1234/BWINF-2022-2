package slice

func MinFunc[T any](s []T, min func(T, T) bool) T {
	minValue := s[0]
	for _, v := range s {
		if min(v, minValue) {
			minValue = v
		}
	}
	return minValue
}
