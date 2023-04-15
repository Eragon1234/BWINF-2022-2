package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils/slice"
)

func FlipAfterBiggest(p pancake.Stack) pancake.SortSteps {
	var sortSteps pancake.SortSteps
	for slice.IndexOfBiggestNonSortedNumber(p) != 0 {
		i := slice.IndexOfBiggestNonSortedNumber(p)
		if i == -1 {
			break
		}
		i = len(p) - i + 1
		sortSteps.Push(int8(i))
		p.Flip(int8(i))

		nsi := slice.NonSortedIndex(p)
		if nsi == -1 {
			break
		}
		nsi = len(p) - nsi
		sortSteps.Push(int8(nsi))
		p.Flip(int8(nsi))
	}
	return sortSteps
}
