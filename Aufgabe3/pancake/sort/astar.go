package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils/slice"
	mySync "BWINF/utils/sync"
	"BWINF/utils/sync/atomic"
	"runtime"
	"strings"
	"sync"
	"time"
)

func BruteForceMultiGoroutineAstar(p pancake.Stack) pancake.SortSteps {
	var wg sync.WaitGroup
	var shortest atomic.Value[string]
	var pq mySync.PriorityQueue[State]

	if !pancake.KeepTrackOfSide {
		// setting the shortest by default to my sort algorithm because it is a possible sort path
		baseShortest := FlipAfterBiggest(*p.Copy())
		shortest.Store(baseShortest.String())
	}
	pq.Push(State{
		Stack: &p,
		Steps: &pancake.SortSteps{},
	}, cost(p))

	run := true

	workerCount := runtime.NumCPU()

	waiting := slice.MakeFunc(workerCount, func(i int) *bool {
		b := false
		return &b
	})
	for i := 0; i < workerCount && run; i++ {
		wg.Add(1)
		go worker(&wg, &run, waiting[i], &pq, &shortest)
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

func worker(wg *sync.WaitGroup, run, waiting *bool, pq *mySync.PriorityQueue[State], shortest *atomic.Value[string]) {
	defer wg.Done()

	pushNew := func(state State) {
		pq.Push(state, cost(*state.Stack))
	}

	pushSolution := func(steps pancake.SortSteps) {
		stepsString := steps.String()
		for s := shortest.Load(); (s == "" || len(steps) < lenOfStepsString(s)) && !shortest.CompareAndSwap(s, stepsString); s = shortest.Load() {
			runtime.Gosched()
		}
	}

	getShortestLength := func() int {
		return lenOfStepsString(shortest.Load())
	}

	for *run {
		*waiting = true
		state, ok := pq.Pop()
		if !ok {
			continue
		}
		*waiting = false
		doState(state, pushNew, pushSolution, getShortestLength)
	}
}

func lenOfStepsString(s string) int {
	return strings.Count(s, " ") + 1
}
