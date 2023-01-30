package pancake

import (
	"Aufgabe3/utils"
	"strconv"
	"strings"
)

type SortSteps[T utils.Number] []T

func (s SortSteps[T]) String() string {
	var stringSteps strings.Builder
	// makes enough space for a single digit and a newline
	stringSteps.Grow(len(s) * 2)
	for _, step := range s {
		stringSteps.WriteString(strconv.Itoa(int(step)))
		stringSteps.WriteString("\n")
	}
	return stringSteps.String()
}

func (s *SortSteps[T]) Push(e T) *SortSteps[T] {
	*s = append(*s, e)
	return s
}

func (s *SortSteps[T]) Copy() *SortSteps[T] {
	newS := make(SortSteps[T], len(*s))
	copy(newS, *s)
	return &newS
}
