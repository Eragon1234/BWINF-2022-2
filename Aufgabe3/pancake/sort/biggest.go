package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/pkg/slice"
)

// FlipAfterBiggest flips at the position after the biggest number and then the biggest number to the bottom
func FlipAfterBiggest(p pancake.Stack) pancake.SortSteps {
	var sortSteps pancake.SortSteps
	for slice.IndexOfBiggestNonSortedNumber(p) != 0 {
		i := slice.IndexOfBiggestNonSortedNumber(p)
		if i == -1 {
			break
		}
		i = len(p) - i + 1
		sortSteps.Push(int8(i))
		p.Flip(i)

		nsi := slice.NonSortedIndex(p)
		if nsi == -1 {
			break
		}
		nsi = len(p) - nsi
		sortSteps.Push(int8(nsi))
		p.Flip(nsi)
	}
	return sortSteps
}
