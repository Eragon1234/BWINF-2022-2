package pancake

import (
	"Aufgabe3/utils"
	"Aufgabe3/utils/slice"
	"strconv"
	"strings"
)

type SortSteps[T utils.Number] []T

func ParseSortSteps[T utils.Number](s string) SortSteps[T] {
	var sortSteps SortSteps[T]
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		step, _ := strconv.Atoi(line)
		sortSteps = append(sortSteps, T(step))
	}
	return sortSteps
}

func (s SortSteps[T]) String() string {
	stringSteps := slice.Map(s, func(e T) string {
		return strconv.Itoa(int(e))
	})
	return strings.Join(stringSteps, "\n")
}

func (s *SortSteps[T]) Push(e T) *SortSteps[T] {
	*s = append(*s, e)
	return s
}

func (s *SortSteps[T]) Copy() *SortSteps[T] {
	newS := make(SortSteps[T], len(*s), cap(*s))
	copy(newS, *s)
	return &newS
}
