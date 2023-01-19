package pancake

import (
	"Aufgabe3/utils"
	"sort"
	"sync"
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

func ShortestBruteForceSortSteps(p Stack) SortSteps {
	sortWays := AllBruteForceSortSteps(p)
	min := sortWays[0]
	for _, sortWay := range sortWays {
		if len(sortWay) < len(min) {
			min = sortWay
		}
	}

	return min
}

func AllBruteForceSortSteps(p Stack) []SortSteps {
	var helper func(*sync.WaitGroup, *utils.TicketSystem[SortSteps], Stack, SortSteps)
	helper = func(wg *sync.WaitGroup, ts *utils.TicketSystem[SortSteps], p Stack, steps SortSteps) {
		defer wg.Done()

		if sort.SliceIsSorted(p, func(i, j int) bool { return p[i] > p[j] }) {
			ts.Put(steps)
			return
		}

		for i := 0; i <= len(p); i++ {
			pancake := p.Copy()
			pancake.Flip(i)

			wg.Add(1)
			go helper(wg, ts, pancake, append(steps, i))
		}
	}

	var wg sync.WaitGroup
	var ts utils.TicketSystem[SortSteps]

	wg.Add(1)
	go helper(&wg, &ts, p, make([]int, 0))

	wg.Wait()

	return ts.GetDone()
}
