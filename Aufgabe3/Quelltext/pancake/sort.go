package pancake

import (
	"Aufgabe3/utils"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
)

func FlipAfterBiggestSortAlgorithm[T utils.Number](p Stack[T]) SortSteps[T] {
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
	var helper func(*sync.WaitGroup, *utils.SyncMap[uint32, SortSteps[T]], Stack[T], SortSteps[T], *atomic.Uint32)
	helper = func(wg *sync.WaitGroup, syncMap *utils.SyncMap[uint32, SortSteps[T]], p Stack[T], steps SortSteps[T], shortest *atomic.Uint32) {
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

		nsi := utils.NonSortedIndex(p)
		wg.Add(len(p) - nsi)
		// running the for loop in reverse because I think that flipping more pancakes has a higher chance of sorting the stack
		// the loop only runs until the non-sorted index because flipping the already sorted pancakes is useless
		for i := len(p); i > nsi; i-- {
			pancake := p.Copy()
			pancake.Flip(i)

			sortSteps := make(SortSteps[T], len(steps))
			copy(sortSteps, steps)
			sortSteps = append(sortSteps, T(i))

			go helper(wg, syncMap, pancake, sortSteps, shortest)
		}
	}

	var wg sync.WaitGroup
	var syncMap utils.SyncMap[uint32, SortSteps[T]]

	var shortest atomic.Uint32
	// setting the shortest number of steps by default to the length of the stack - 1
	// because only one element would be left and the whole stack would be sorted
	shortest.Store(uint32(len(p) - 1))

	wg.Add(1)
	go helper(&wg, &syncMap, p, SortSteps[T]{}, &shortest)

	wg.Wait()

	value, ok := syncMap.Load(shortest.Load())
	if ok {
		return value
	}
	return nil
}
