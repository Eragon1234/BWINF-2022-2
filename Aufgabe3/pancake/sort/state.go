package sort

import "BWINF/Aufgabe3/pancake"

// State is a struct that contains a stack and a list of steps.
// It is used to store the current state of the sorting process.
type State struct {
	Stack *pancake.Stack
	Steps *pancake.SortSteps
}
