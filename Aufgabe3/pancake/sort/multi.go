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

func BruteForceMultiGoroutine(p pancake.Stack) pancake.SortSteps {
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
		time.Sleep(time.Millisecond * 500)
		runtime.Gosched()
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

	for *run {
		*waiting = true
		state, ok := pq.Pop()
		if !ok {
			continue
		}
		*waiting = false
		doStack(state, pq, shortest)
	}
}

func doStack(item State, pq *mySync.PriorityQueue[State], shortest *atomic.Value[string]) {
	p := *item.Stack
	steps := *item.Steps

	lenOfSteps := len(steps)

	// check current steps length is greater than or equal to the smallest steps in done
	if s := shortest.Load(); s != "" && lenOfSteps >= lenOfStepsString(s) {
		return
	}

	nonSortedIndex := slice.NonSortedIndex(p)
	var negativeCount int
	if pancake.KeepTrackOfSide {
		negativeCount = slice.CountFunc(p, func(i int8) bool { return i < 0 })
	}
	// when sorted index is -1 the stack is sorted
	if nonSortedIndex == -1 && negativeCount == 0 {
		stepsString := steps.String()
		for s := shortest.Load(); s == "" || lenOfSteps < lenOfStepsString(s) && !shortest.CompareAndSwap(s, stepsString); s = shortest.Load() {
			runtime.Gosched()
		}
		return
	}

	// check if the next iteration won't exit early because the current steps length is greater than or equal to the smallest steps in done
	// exit early if the next iteration will exit early, to prevent the spawning unnecessary goroutines
	if s := shortest.Load(); s != "" && lenOfSteps+1 >= lenOfStepsString(s) {
		return
	}

	// updating the stack to only contain the unsorted pancakes because we can ignore the sorted ones
	// it won't affect the indexes because we are counting the flip index from the top of the stack
	p = p[nonSortedIndex:]

	// running the for loop in reverse because I think that flipping more pancakes has a higher chance of sorting the stack
	for i := len(p); i >= 0; i-- {
		newP := p.Copy().Flip(int8(i))
		pq.Push(State{
			Stack: newP,
			Steps: steps.Copy().Push(int8(i)),
		}, cost(*newP))
	}
}

func cost(p pancake.Stack) uint8 {
	if len(p) == 0 {
		return 0
	}
	if len(p) < 3 {
		return uint8(len(p))
	}
	var count uint8 = 1
	reducing := p[0] > p[1]
	for i := 1; i < len(p)-1; i++ {
		if p[i] > p[i+1] != reducing {
			count++
			if i+2 < len(p) {
				reducing = p[i+1] > p[i+2]
			}
		}
	}
	return count + uint8(len(p))
}

func lenOfStepsString(s string) int {
	return strings.Count(s, " ") + 1
}
