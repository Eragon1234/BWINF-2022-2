package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils/slice"
)

// doState does the algorithm for a given state
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
		negativeCount = slice.CountFunc(p, func(i int8) bool { return i < 0 })
	}

	if nonSortedIndex == -1 && negativeCount == 0 {
		pushSolution(steps)
		return
	}

	if lenOfSteps-1 >= getShortestLength() {
		return
	}

	for i := len(p); i >= 0; i-- {
		var newStack pancake.Stack
		if i == 0 {
			newStack = p
		} else {
			newStack = *p.Copy()
		}
		pushNew(State{
			Stack: newStack.Flip(i),
			Steps: steps.Copy().Push(int8(i)),
		})
	}
}
