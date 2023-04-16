package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/pkg"
	"BWINF/pkg/slice"
)

// CalculatePWUE calculates PWUE for a pancake of size n.
func CalculatePWUE(n int) (stack pancake.Stack, sortSteps pancake.SortSteps) {
	if n == 0 {
		return pancake.Stack{}, pancake.SortSteps{}
	}
	nums := slice.MakeFunc(n, func(i int) int {
		return i + 1
	})
	perm := pkg.NewPermutation(nums)
	for p := nums; p != nil; p = perm.Next() {
		s := slice.Map(p, func(i int) int8 {
			return int8(i)
		})
		newSortSteps := BruteForce(s)
		if len(newSortSteps) > len(sortSteps) {
			sortSteps = newSortSteps
			stack = s
		}
	}
	return stack, sortSteps
}
