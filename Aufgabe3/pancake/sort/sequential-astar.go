package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/pkg/queue"
	"BWINF/pkg/set"
	"math"
)

// sequentialAstar is a single goroutine implementation of A* to find the shortest path to sort a pancake stack
func sequentialAstar(p pancake.Stack) pancake.SortSteps {
	var shortest pancake.SortSteps
	pq := queue.PriorityQueue[State]{}
	pq.Push(State{
		Stack: &p,
		Steps: &pancake.SortSteps{},
	}, 0)

	var visited set.Set[string]

	if !pancake.KeepTrackOfSide {
		// setting the shortest by default to my sort algorithm because it is a possible sort path
		baseShortest := FlipAfterBiggest(*p.Copy())
		shortest = baseShortest
	}

	pushNew := func(state State) {
		stateString := state.Stack.String()
		if visited.Contains(stateString) {
			return
		}
		visited.Add(stateString)
		pq.Push(state, cost(state))
	}

	pushSolution := func(steps pancake.SortSteps) {
		if shortest == nil || len(steps) < len(shortest) {
			shortest = steps
		}
	}

	getShortestLength := func() int {
		if shortest == nil {
			return math.MaxInt
		}
		return len(shortest)
	}

	for pq.Len() != 0 {
		state, _ := pq.Pop() // won't be empty
		doState(state, pushNew, pushSolution, getShortestLength)
	}

	return shortest
}
