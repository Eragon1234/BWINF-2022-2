package utils

import "math/rand"

// ChooseWeighted chooses an index from a slice of weights.
// the weights are relative to each other
func ChooseWeighted(weights []float64) int {
	total := 0.0
	for _, w := range weights {
		total += w
	}
	r := rand.Float64() * total
	for i, w := range weights {
		if r < w {
			return i
		}
		r -= w
	}
	return len(weights) - 1
}
