package pkg

// Min returns the smaller of the two values.
// If a and b are equal, b is returned.
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
