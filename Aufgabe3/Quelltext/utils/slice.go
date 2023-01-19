package utils

func IndexOfBiggestInt(s []int) int {
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

func NonSortedIndex(s []int) int {
	for i := range s {
		if IndexOfBiggestInt(s[i:]) != 0 {
			return i
		}
	}
	return -1
}

func IndexOfBiggestNonSortedInt(s []int) int {
	nsi := NonSortedIndex(s)
	if nsi == -1 {
		return -1
	}
	return IndexOfBiggestInt(s[nsi+1:]) + nsi + 1
}
