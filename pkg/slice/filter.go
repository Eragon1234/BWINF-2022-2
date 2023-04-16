package slice

// FilterFunc filters the slice s with the filter function filter.
// The filter function should return true if the element should be filtered.
func FilterFunc[T any](s []T, filter func(T) bool) []T {
	var result []T
	for _, v := range s {
		if !filter(v) {
			result = append(result, v)
		}
	}
	return result
}
