package pancake

import (
	"sort"
)

func ShortestBruteForceSortSteps(p Pancake) []int {
	sortWays := AllBruteForceSortSteps(p)
	min := sortWays[0]
	for _, sortWay := range sortWays {
		if len(sortWay) < len(min) {
			min = sortWay
		}
	}

	return min
}

func AllBruteForceSortSteps(p Pancake) [][]int {
	var helper func(Pancake, []int) [][]int
	helper = func(p Pancake, steps []int) [][]int {
		if sort.SliceIsSorted(p, func(i, j int) bool { return p[i] > p[j] }) {
			return [][]int{steps}
		}

		sortWays := make([][]int, 0)
		for i := 0; i <= len(p); i++ {
			pancake := p.Copy()
			pancake.Flip(i)
			sortWay := helper(pancake, append(steps, i))
			sortWays = append(sortWays, sortWay...)
		}

		return sortWays
	}

	return helper(p, make([]int, 0))
}
