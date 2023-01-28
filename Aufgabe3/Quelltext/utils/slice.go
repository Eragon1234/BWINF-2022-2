package utils

func IndexOfBiggest[T Number](s []T) int { // O(n)
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

func NonSortedIndex[T Number](s []T) int { // O(n^2)
	for i := range s {
		if IndexOfBiggest(s[i:]) != 0 {
			return i
		}
	}
	return -1
}

func IndexOfBiggestNonSortedNumber[T Number](s []T) int { // O(n^2)
	nsi := NonSortedIndex(s)
	if nsi == -1 {
		return -1
	}
	return IndexOfBiggest(s[nsi+1:]) + nsi + 1
}
