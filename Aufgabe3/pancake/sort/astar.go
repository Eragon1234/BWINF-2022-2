package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/pkg/slice"
	mySync "BWINF/pkg/sync"
	"BWINF/pkg/sync/atomic"
	"math"
	"runtime"
	"sync"
	"time"
)

// BruteForceMultiGoroutineAstar is a brute force algorithm that uses multiple goroutines to find the shortest sort path using the A* algorithm
func BruteForceMultiGoroutineAstar(p pancake.Stack) pancake.SortSteps {
	var wg sync.WaitGroup
	var shortest atomic.Value[string]
	var pq mySync.PriorityQueue[State]
	var visited mySync.Set[string]

	if !pancake.KeepTrackOfSide {
		// setting the shortest by default to my sort algorithm because it is a possible sort path
		baseShortest := FlipAfterBiggest(*p.Copy())
		shortest.Store(baseShortest.String())
	}
	pq.Push(State{
		Stack: &p,
		Steps: &pancake.SortSteps{},
	}, cost(p))

	getNext := func() (State, bool) {
		return pq.Pop()
	}

	pushNew := func(state State) {
		stackString := state.Stack.String()
		if visited.Contains(stackString) {
			return
		}
		visited.Add(stackString)
		pq.Push(state, cost(*state.Stack))
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

	run := true

	workerCount := runtime.NumCPU()

	waiting := slice.MakeFunc(workerCount, func(i int) *bool {
		b := false
		return &b
	})
	for i := 0; i < workerCount && run; i++ {
		wg.Add(1)
		go worker(&wg, &run, waiting[i], getNext, pushNew, pushSolution, getShortestLength)
	}

	for slice.CountFunc(waiting, func(b *bool) bool { return *b }) != len(waiting) || pq.Len() != 0 {
		runtime.Gosched()
		time.Sleep(time.Millisecond * 500)
	}

	run = false

	wg.Wait()

	value := shortest.Load()
	if value == "" {
		return pancake.SortSteps{}
	}

	sortSteps := pancake.ParseSortSteps(value)

	return sortSteps
}

func worker(wg *sync.WaitGroup, run, waiting *bool, getNext func() (State, bool), pushNew func(State), pushSolution func(pancake.SortSteps), getShortestLength func() int) {
	defer wg.Done()

	for *run {
		*waiting = true
		state, ok := getNext()
		if !ok {
			continue
		}
		*waiting = false
		doState(state, pushNew, pushSolution, getShortestLength)
	}
}
