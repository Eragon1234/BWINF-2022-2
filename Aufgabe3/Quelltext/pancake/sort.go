package pancake

import (
	"Aufgabe3/utils"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func FlipAfterBiggestSortAlgorithm[T utils.Number](p Stack[T]) SortSteps[T] { // nearly works
	var sortSteps SortSteps[T]
	for utils.IndexOfBiggestNonSortedInt(p) != 0 {
		i := utils.IndexOfBiggestNonSortedInt(p)
		if i == -1 {
			break
		}
		i = len(p) - i + 1
		sortSteps = append(sortSteps, T(i))
		p.Flip(i)

		nsi := utils.NonSortedIndex(p)
		if nsi == -1 {
			break
		}
		nsi = len(p) - nsi
		sortSteps = append(sortSteps, T(nsi))
		p.Flip(nsi)
	}
	return sortSteps
}

func BruteForceSort[T utils.Number](p Stack[T]) SortSteps[T] {
	var helper func(*sync.WaitGroup, *utils.AtomicValue[string], Stack[T], SortSteps[T], int)
	helper = func(wg *sync.WaitGroup, shortest *utils.AtomicValue[string], p Stack[T], steps SortSteps[T], maxSteps int) {
		defer wg.Done()

		lenOfSteps := len(steps)
		sortedIndex := utils.NonSortedIndex(p)

		// check current steps length is greater than or equal to the smallest steps in done
		if s, ok := shortest.Load(); ok && lenOfSteps >= utils.Min(lenOfStepsString(s), maxSteps) {
			return
		}

		// when sorted index is -1 the stack is sorted
		if sortedIndex == -1 {
			for s, ok := shortest.Load(); (!ok || lenOfSteps < lenOfStepsString(s)) && !shortest.CompareAndSwap(s, steps.String()); s, ok = shortest.Load() {
				if !ok {
					shortest.Store(steps.String())
					return
				}
				runtime.Gosched()
			}

			return
		}

		if s, ok := shortest.Load(); ok && lenOfSteps >= utils.Min(lenOfStepsString(s), maxSteps)+1 {
			return
		}

		relevantP := p[sortedIndex:]
		// running the for loop in reverse because I think that flipping more pancakes has a higher chance of sorting the stack
		wg.Add(len(relevantP))
		for i := len(relevantP); i > 0; i-- {
			go helper(wg, shortest, *relevantP.Copy().Flip(i), *steps.Copy().Push(T(i)), maxSteps)
		}
	}

	var wg sync.WaitGroup
	var shortest utils.AtomicValue[string]

	wg.Add(1)
	go helper(&wg, &shortest, p, SortSteps[T]{}, len(p)-1)

	wg.Wait()

	value, ok := shortest.Load()
	if !ok {
		panic("no path found")
	}

	var sortSteps SortSteps[T]
	for _, line := range strings.Split(value, "\n") {
		if line == "" {
			continue
		}
		step, _ := strconv.Atoi(line)
		sortSteps = append(sortSteps, T(step))
	}

	return sortSteps
}

func lenOfStepsString(steps string) int {
	// because every step is followed by a newline character we can count the number of new line characters to get the number of steps
	return strings.Count(steps, "\n")
}
