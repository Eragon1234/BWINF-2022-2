package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils/queue"
	"BWINF/utils/slice"
)

func BruteForce(p pancake.Stack) pancake.SortSteps {
	var shortest pancake.SortSteps
	pq := queue.PriorityQueue[State]{}
	pq.Push(State{
		Stack: &p,
		Steps: &pancake.SortSteps{},
	}, 0)

	if !pancake.KeepTrackOfSide {
		// setting the shortest by default to my sort algorithm because it is a possible sort path
		baseShortest := FlipAfterBiggest(*p.Copy())
		shortest = baseShortest
	}

	for pq.Len() != 0 {
		state, _ := pq.Pop() // won't be empty
		if shortest != nil && len(*state.Steps) >= len(shortest) {
			continue
		}

		nonSortedIndex := slice.NonSortedIndex(p)
		var negativeCount int
		if pancake.KeepTrackOfSide {
			negativeCount = slice.CountFunc(*state.Stack, func(i int8) bool { return i < 0 })
		}

		if nonSortedIndex == -1 && negativeCount == 0 {
			shortest = *state.Steps
			continue
		}

		for i := len(*state.Stack); i >= 0; i-- {
			newStack := *state.Stack.Copy()
			newStack.Flip(int8(i))
			newSteps := *state.Steps.Copy()
			newSteps.Push(int8(i))
			pq.Push(State{
				Stack: &newStack,
				Steps: &newSteps,
			}, uint8(len(newSteps)))
		}
	}

	return shortest
}
