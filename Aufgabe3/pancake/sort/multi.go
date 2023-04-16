package sort

import (
	"BWINF/Aufgabe3/pancake"
	mySync "BWINF/pkg/sync"
	"BWINF/pkg/sync/atomic"
	"math"
	"runtime"
	"sync"
)

// BruteForceMultiGoroutine brute forces the sorting of a pancake stack with multiple goroutines
func BruteForceMultiGoroutine(p pancake.Stack) pancake.SortSteps {
	var helper func(*sync.WaitGroup, State)
	var wg sync.WaitGroup
	var shortest atomic.Value[string]
	var visited mySync.Set[string]

	pushNew := func(state State) {
		stackString := state.Stack.String()
		if visited.Contains(stackString) {
			return
		}
		visited.Add(stackString)
		wg.Add(1)
		go helper(&wg, state)
	}

	pushSolution := func(steps pancake.SortSteps) {
		stepsString := steps.String()
		for s := shortest.Load(); (s == "" || len(steps) < pancake.LenOfSortStepsString(s)) && !shortest.CompareAndSwap(s, stepsString); s = shortest.Load() {
			runtime.Gosched()
		}
	}

	getShortestLength := func() int {
		s := shortest.Load()
		if s == "" {
			return math.MaxInt
		}
		return pancake.LenOfSortStepsString(s)
	}

	helper = func(wg *sync.WaitGroup, state State) {
		defer wg.Done()

		doState(state, pushNew, pushSolution, getShortestLength)
	}

	baseShortest := make(pancake.SortSteps, 0)
	// setting the shortest by default to my sort algorithm because it is a possible sort path
	if !pancake.KeepTrackOfSide {
		baseShortest = FlipAfterBiggest(*p.Copy())
		shortest.Store(baseShortest.String())
	}

	var steps = make(pancake.SortSteps, 0, len(baseShortest))
	wg.Add(1)
	go helper(&wg, State{
		Stack: &p,
		Steps: &steps,
	})

	wg.Wait()

	value := shortest.Load()
	if value == "" {
		return pancake.SortSteps{}
	}

	sortSteps := pancake.ParseSortSteps(value)

	return sortSteps
}
