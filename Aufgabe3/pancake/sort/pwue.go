package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils"
	"BWINF/utils/slice"
)

func CalculatePWUE(n int) (stack pancake.Stack, sortSteps pancake.SortSteps) {
	if n == 0 {
		return pancake.Stack{}, pancake.SortSteps{}
	}
	nums := slice.MakeFunc(n, func(i int) int {
		return i + 1
	})
	perm := utils.NewPermutation(nums)
	for p := nums; p != nil; p = perm.Next() {
		s := slice.Map(p, func(i int) int8 {
			return int8(i)
		})
		newSortSteps := BruteForceSort[int8](s)
		if len(newSortSteps) > len(sortSteps) {
			sortSteps = newSortSteps
			stack = s
		}
	}
	return stack, sortSteps
}
