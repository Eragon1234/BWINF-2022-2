package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils/slice"
)

func doState(state State, pushNew func(State), pushSolution func(steps pancake.SortSteps), getShortestLength func() int) {
	p := *state.Stack
	steps := *state.Steps

	lenOfSteps := len(steps)

	if lenOfSteps >= getShortestLength() {
		return
	}

	nonSortedIndex := slice.NonSortedIndex(p)
	var negativeCount int
	if pancake.KeepTrackOfSide {
		negativeCount = slice.CountFunc(p, func(i int) bool { return i < 0 })
	}

	if nonSortedIndex == -1 && negativeCount == 0 {
		pushSolution(steps)
		return
	}

	if negativeCount == 0 {
		p = p[nonSortedIndex:]
	}

	if lenOfSteps-1 >= getShortestLength() {
		return
	}

	for i := len(p); i >= 0; i-- {
		pushNew(State{
			Stack: p.Copy().Flip(i),
			Steps: steps.Copy().Push(i),
		})
	}
}
