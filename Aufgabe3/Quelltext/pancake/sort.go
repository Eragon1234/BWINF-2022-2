package pancake

import (
	"Aufgabe3/utils"
	"sort"
	"sync"
)

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
