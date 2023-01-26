package pancake

import (
	"Aufgabe3/utils"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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
	var helper func(*sync.WaitGroup, *atomic.Value, Stack[T], SortSteps[T], int)
	helper = func(wg *sync.WaitGroup, shortest *atomic.Value, p Stack[T], steps SortSteps[T], maxSteps int) {
		defer wg.Done()

		lenOfSteps := len(steps)
		sortedIndex := utils.NonSortedIndex(p)

		// check current steps length is greater than or equal to the smallest steps in done
		if s := shortest.Load(); s != nil && lenOfSteps >= utils.Min(int(math.Floor(float64(len(s.(string)))/2)), maxSteps) {
			return
		}

		if sortedIndex == -1 {
			var stringSteps strings.Builder
			stringSteps.Grow(len(steps))
			for _, step := range steps {
				stringSteps.WriteString(strconv.Itoa(int(step)))
				stringSteps.WriteString("\n")
			}
			for s := shortest.Load(); (s == nil || lenOfSteps < int(math.Floor(float64(len(s.(string)))/2))) && !shortest.CompareAndSwap(s, stringSteps.String()); s = shortest.Load() {
				runtime.Gosched()
			}

			return
		}

		if s := shortest.Load(); s != nil && lenOfSteps >= utils.Min(int(math.Floor(float64(len(s.(string)))/2)), maxSteps)+1 {
			return
		}

		relevantP := p[sortedIndex:]
		// running the for loop in reverse because I think that flipping more pancakes has a higher chance of sorting the stack
		wg.Add(len(relevantP))
		for i := len(relevantP); i > 0; i-- {
			pancake := relevantP.Copy()
			pancake.Flip(i)

			sortSteps := make(SortSteps[T], len(steps))
			copy(sortSteps, steps)
			sortSteps = append(sortSteps, T(i))

			go helper(wg, shortest, pancake, sortSteps, maxSteps)
		}
	}

	var wg sync.WaitGroup
	var shortest atomic.Value

	wg.Add(1)
	go helper(&wg, &shortest, p, SortSteps[T]{}, len(p)-1)

	wg.Wait()

	value := shortest.Load()

	var sortSteps SortSteps[T]
	for _, line := range strings.Split(value.(string), "\n") {
		if line == "" {
			continue
		}
		step, _ := strconv.Atoi(line)
		sortSteps = append(sortSteps, T(step))
	}

	return sortSteps
}
