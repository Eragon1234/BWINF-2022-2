package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils/slice"
	mySync "BWINF/utils/sync"
	"BWINF/utils/sync/atomic"
	"runtime"
	"strings"
	"sync"
)

func BruteForceMultiGoroutineInline(p pancake.Stack) pancake.SortSteps {
	var helper func(*sync.WaitGroup, *atomic.Value[string], *mySync.Set[string], pancake.Stack, pancake.SortSteps)
	helper = func(wg *sync.WaitGroup, shortest *atomic.Value[string], visited *mySync.Set[string], p pancake.Stack, steps pancake.SortSteps) {
		defer wg.Done()

		lenOfSteps := len(steps)

		// check current steps length is greater than or equal to the smallest steps in done
		if s := shortest.Load(); s != "" && lenOfSteps >= lenOfStepsString(s) {
			return
		}

		nonSortedIndex := slice.NonSortedIndex(p)
		var negativeCount int
		if pancake.KeepTrackOfSide {
			negativeCount = slice.CountFunc(p, func(e int8) bool {
				return e < 0
			})
		}

		// when sorted index is -1 the stack is sorted
		if nonSortedIndex == -1 && negativeCount == 0 {
			stepsString := steps.String()
			for s := shortest.Load(); (s == "" || lenOfSteps < lenOfStepsString(s)) && !shortest.CompareAndSwap(s, stepsString); s = shortest.Load() {
				runtime.Gosched()
			}
			return
		}

		// check if the next iteration won't exit early because the current steps length is greater than or equal to the smallest steps in done
		// exit early if the next iteration will exit early, to prevent the spawning unnecessary goroutines
		if s := shortest.Load(); s != "" && lenOfSteps+1 >= lenOfStepsString(s) {
			return
		}

		//if negativeCount == 0 {
		//	// updating the stack to only contain the unsorted pancakes because we can ignore the sorted ones,
		//	// it won't affect the indexes because we are counting the flip index from the top of the stack
		//	p = p[nonSortedIndex:]
		//}

		// running the for loop in reverse because I think that flipping more pancakes has a higher chance of sorting the stack
		for i := len(p); i > 0; i-- {
			var newP pancake.Stack
			if i == 1 {
				newP = *p.Flip(1)
			} else {
				newP = *p.Copy().Flip(i)
			}
			newPString := newP.String()

			// check if the new stack is already visited
			// we can't reach the same state with a longer path because
			if visited.Contains(newPString) {
				continue
			}
			visited.Add(newPString)

			wg.Add(1)
			go helper(wg, shortest, visited, newP, *steps.Copy().Push(int8(i)))
		}
	}

	var wg sync.WaitGroup
	var shortest atomic.Value[string]

	baseShortest := make(pancake.SortSteps, 0, len(p))
	// setting the shortest by default to my sort algorithm because it is a possible sort path
	if !pancake.KeepTrackOfSide {
		baseShortest = FlipAfterBiggest(*p.Copy())
		shortest.Store(baseShortest.String())
	}

	wg.Add(1)
	// creating the sort steps with capacity of the length of the base shortest because we won't need more than that
	// the higher capacity is to prevent the slice from being reallocated when we append to it which leads to performance issues
	// the only problem is that we won't always use the full capacity which increases the memory usage
	// preventing the reallocation also prevent heap fragmentation
	go helper(&wg, &shortest, new(mySync.Set[string]), p, make(pancake.SortSteps, 0, len(baseShortest)))

	wg.Wait()

	value := shortest.Load()
	if value == "" {
		return pancake.SortSteps{}
	}

	sortSteps := pancake.ParseSortSteps(value)

	return sortSteps
}

func lenOfStepsString(steps string) int {
	return strings.Count(steps, " ") + 1
}
