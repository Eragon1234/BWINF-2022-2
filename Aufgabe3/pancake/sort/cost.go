package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/utils/slice"
)

// cost returns the cost of a given pancake stack.
func cost(p pancake.Stack) int {
	negativeCount := slice.CountFunc(p, func(i int8) bool { return i < 0 })
	if len(p) < 3 {
		return len(p) + negativeCount
	}
	var count = 1
	reducing := p[0] > p[1]
	for i := 1; i < len(p)-1; i++ {
		if p[i] > p[i+1] != reducing {
			count++
			if i+2 < len(p) {
				reducing = p[i+1] > p[i+2]
			}
		}
	}
	return count + len(p) + negativeCount
}
