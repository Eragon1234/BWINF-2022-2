package pancake

import (
	"BWINF/utils"
	"BWINF/utils/slice"
	"BWINF/utils/sync/atomic"
	"runtime"
	"strings"
	"sync"
)

func FlipAfterBiggestSortAlgorithm[T utils.Number](p Stack[T]) SortSteps[T] { // nearly works
	var sortSteps SortSteps[T]
	for slice.IndexOfBiggestNonSortedNumber(p) != 0 {
		i := slice.IndexOfBiggestNonSortedNumber(p)
		if i == -1 {
			break
		}
		i = len(p) - i + 1
		sortSteps.Push(T(i))
		p.Flip(i)

		nsi := slice.NonSortedIndex(p)
		if nsi == -1 {
			break
		}
		nsi = len(p) - nsi
		sortSteps.Push(T(nsi))
		p.Flip(nsi)
	}
	return sortSteps
}

func BruteForceSort[T utils.Number](p Stack[T]) SortSteps[T] {
	var helper func(*sync.WaitGroup, *atomic.Value[string], Stack[T], SortSteps[T])
	helper = func(wg *sync.WaitGroup, shortest *atomic.Value[string], p Stack[T], steps SortSteps[T]) {
		defer wg.Done()

		lenOfSteps := len(steps)

		// check current steps length is greater than or equal to the smallest steps in done
		if s := shortest.Load(); s != "" && lenOfSteps >= lenOfStepsString(s) {
			return
		}

		nonSortedIndex := slice.NonSortedIndex(p)

		// when sorted index is -1 the stack is sorted
		if nonSortedIndex == -1 {
			stepsString := steps.String()
			for s := shortest.Load(); s != "" && lenOfSteps < lenOfStepsString(s) && !shortest.CompareAndSwap(s, stepsString); s = shortest.Load() {
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

		wg.Add(len(p))
		// running the for loop in reverse because I think that flipping more pancakes has a higher chance of sorting the stack
		for i := len(p); i > 0; i-- {
			go helper(wg, shortest, *p.Copy().Flip(i), *steps.Copy().Push(T(i)))
		}
	}

	var wg sync.WaitGroup
	var shortest atomic.Value[string]

	// setting the shortest by default to my sort algorithm because it is a possible sort path
	baseShortest := FlipAfterBiggestSortAlgorithm(*p.Copy())
	shortest.Store(baseShortest.String())

	wg.Add(1)
	// creating the sort steps with capacity of the length of the base shortest - 1 because we won't need more than that
	// the higher capacity is to prevent the slice from being reallocated when we append to it which leads to performance issues
	// the only problem is that we won't always use the full capacity which increases the memory usage
	// preventing the reallocation also prevent heap fragmentation
	go helper(&wg, &shortest, p, make(SortSteps[T], 0, len(baseShortest)-1))

	wg.Wait()

	value := shortest.Load()
	if value == "" {
		return SortSteps[T]{}
	}

	sortSteps := ParseSortSteps[T](value)

	return sortSteps
}

func lenOfStepsString(steps string) int {
	// because every step is followed by a newline character we can count the number of new line characters to get the number of steps
	return strings.Count(steps, " ")
}
