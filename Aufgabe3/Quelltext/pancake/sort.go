package pancake

import (
	"Aufgabe3/utils"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
)

func FlipAfterBiggestSortAlgorithm(p Stack) SortSteps {
	var sortSteps SortSteps
	for utils.IndexOfBiggestNonSortedInt(p) != 0 {
		i := utils.IndexOfBiggestNonSortedInt(p)
		if i == -1 {
			break
		}
		i = len(p) - i + 1
		sortSteps = append(sortSteps, i)
		p.Flip(i)

		nsi := utils.NonSortedIndex(p)
		if nsi == -1 {
			break
		}
		nsi = len(p) - nsi
		sortSteps = append(sortSteps, nsi)
		p.Flip(nsi)
	}
	return sortSteps
}

func BruteForceSort(p Stack) SortSteps {
	var helper func(*sync.WaitGroup, *utils.SyncMap[uint32, SortSteps], Stack, SortSteps, *atomic.Uint32)
	helper = func(wg *sync.WaitGroup, syncMap *utils.SyncMap[uint32, SortSteps], p Stack, steps SortSteps, shortest *atomic.Uint32) {
		defer wg.Done()

		// check current steps length is greater than or equal to the smallest steps in done
		if uint32(len(steps)) >= shortest.Load() {
			return
		}

		if sort.SliceIsSorted(p, func(i, j int) bool { return p[i] > p[j] }) {
			l := uint32(len(steps))

			for s := shortest.Load(); l < s && !shortest.CompareAndSwap(s, l); s = shortest.Load() {
				runtime.Gosched()
			}

			syncMap.Store(l, steps)
			return
		}

		if uint32(len(steps)) >= shortest.Load()+1 {
			return
		}

		wg.Add(len(p))
		// running the for loop in reverse because I think that flipping more pancakes has a higher chance of sorting the stack
		// the loop doesn't include 0 because it has the same as flipping 1
		for i := len(p); i > 0; i-- {
			pancake := p.Copy()
			pancake.Flip(i)

			go helper(wg, syncMap, pancake, append(SortSteps{i}, steps...), shortest)
		}
	}

	var wg sync.WaitGroup
	var syncMap utils.SyncMap[uint32, SortSteps]

	var shortest atomic.Uint32
	// setting the shortest number of steps by default to the length of the stack - 1
	// because only one element would be left and the whole stack would be sorted
	shortest.Store(uint32(len(p) - 1))

	wg.Add(1)
	go helper(&wg, &syncMap, p, SortSteps{}, &shortest)

	wg.Wait()

	value, ok := syncMap.Load(shortest.Load())
	if ok {
		return value
	}
	return nil
}
