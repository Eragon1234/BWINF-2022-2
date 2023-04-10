package slice

func FilterFunc[T any](s []T, filter func(T) bool) []T {
	var result []T
	for _, v := range s {
		if !filter(v) {
			result = append(result, v)
		}
	}
	return result
}
