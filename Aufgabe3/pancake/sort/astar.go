package sort

import (
	"BWINF/Aufgabe3/pancake"
	"runtime"
)

var WorkerCount = runtime.NumCPU()

// Astar calls sequentialAstar if WorkerCount is 1 otherwise it calls concurrentAstar
func Astar(p pancake.Stack) pancake.SortSteps {
	if WorkerCount == 1 {
		return sequentialAstar(p)
	}
	return concurrentAstar(p)
}
