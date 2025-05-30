package slice

// ReverseSlice reverses the order of the elements in the slice.
func ReverseSlice[S any](s []S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
