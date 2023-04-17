package sort

import (
	"BWINF/Aufgabe3/pancake"
)

var WorkerCount = 1

// Astar calls sequentialAstar if WorkerCount is 1 otherwise it calls concurrentAstar
func Astar(p pancake.Stack) pancake.SortSteps {
	if WorkerCount == 1 {
		return sequentialAstar(p)
	}
	return concurrentAstar(p)
}
