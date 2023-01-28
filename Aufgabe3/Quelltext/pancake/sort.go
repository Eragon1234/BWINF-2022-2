package pancake

import (
	"Aufgabe3/utils"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func FlipAfterBiggestSortAlgorithm[T utils.Number](p Stack[T]) SortSteps[T] { // nearly works
	var sortSteps SortSteps[T]
	for utils.IndexOfBiggestNonSortedNumber(p) != 0 {
		i := utils.IndexOfBiggestNonSortedNumber(p)
		if i == -1 {
			break
		}
		i = len(p) - i + 1
		sortSteps.Push(T(i))
		p.Flip(i)

		nsi := utils.NonSortedIndex(p)
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
	var helper func(*sync.WaitGroup, *utils.AtomicValue[string], Stack[T], SortSteps[T], int)
	helper = func(wg *sync.WaitGroup, shortest *utils.AtomicValue[string], p Stack[T], steps SortSteps[T], maxSteps int) {
		defer wg.Done()

		lenOfSteps := len(steps)

		// check current steps length is greater than or equal to the smallest steps in done
		if s, ok := shortest.Load(); ok && lenOfSteps >= utils.Min(lenOfStepsString(s), maxSteps) {
			return
		}

		nonSortedIndex := utils.NonSortedIndex(p)

		// when sorted index is -1 the stack is sorted
		if nonSortedIndex == -1 {
			stepsString := steps.String()
			for s, ok := shortest.Load(); (!ok || lenOfSteps < lenOfStepsString(s)) && !shortest.CompareAndSwap(s, stepsString); s, ok = shortest.Load() {
				runtime.Gosched()
			}
			return
		}

		// updating the stack to only contain the unsorted pancakes because we can ignore the sorted ones
		p = p[nonSortedIndex:]

		wg.Add(len(p))
		// running the for loop in reverse because I think that flipping more pancakes has a higher chance of sorting the stack
		for i := len(p); i > 0; i-- {
			go helper(wg, shortest, *p.Copy().Flip(i), *steps.Copy().Push(T(i)), maxSteps)
		}
	}

	var wg sync.WaitGroup
	var shortest utils.AtomicValue[string]

	wg.Add(1)
	go helper(&wg, &shortest, p, SortSteps[T]{}, len(p)-1)

	wg.Wait()

	value, ok := shortest.Load()
	if !ok {
		log.Fatalln("no path found")
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
