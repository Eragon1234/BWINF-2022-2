package sort

// cost returns the cost of a given pancake stack.
func cost(state State) int {
	return len(*state.Steps)
}
