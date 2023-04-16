package pkg

// Min returns the smaller of the two values.
// If a and b are equal, b is returned.
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of the two values.
// If a and b are equal, b is returned.
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Clamp returns the value clamped between min and max.
func Clamp[T Number](min, value, max T) T {
	return Min[T](max, Max[T](min, value))
}
