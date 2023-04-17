package slice

import (
	"BWINF/pkg"
)

// IndexOfBiggest returns the index of the biggest number in the slice.
func IndexOfBiggest[T pkg.Number](s []T) int { // O(n)
	if len(s) == 0 {
		return -1
	}

	maxIndex := 0
	for i := range s {
		if s[i] > s[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}

// NonSortedIndex returns the index of the first number that is not sorted.
// If the slice is sorted, -1 is returned.
func NonSortedIndex[T pkg.Number](s []T) int { // O(n)
	for i := range s {
		if IndexOfBiggest(s[i:]) != 0 {
			return i
		}
	}
	return -1
}

// IndexOfBiggestNonSortedNumber returns the index of the biggest number that is not sorted.
func IndexOfBiggestNonSortedNumber[T pkg.Number](s []T) int { // O(n^2)
	nsi := NonSortedIndex(s)
	if nsi == -1 {
		return -1
	}
	return IndexOfBiggest(s[nsi+1:]) + nsi + 1
}
