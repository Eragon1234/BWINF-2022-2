package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils/slice"
)

func CalculatePWUE(n int) (stack pancake.Stack[int8], sortSteps pancake.SortSteps[int8]) {
	if n == 0 {
		return pancake.Stack[int8]{}, pancake.SortSteps[int8]{}
	}
	nums := slice.MakeFunc(n, func(i int) int {
		return i + 1
	})
	perm := NewPermutation(nums)
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

type Permutation struct {
	original    []int
	permutation []int
}

func NewPermutation(s []int) Permutation {
	return Permutation{
		original:    s,
		permutation: make([]int, len(s)),
	}
}

// Next returns the next permutation or nil if there are no more permutations.
// The permutations don't include the original.
func (p *Permutation) Next() []int {
	nextPerm(p.permutation)
	ok := p.permutation[0] < len(p.permutation)
	if !ok {
		return nil
	}
	return getPerm(p.original, p.permutation)
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}
