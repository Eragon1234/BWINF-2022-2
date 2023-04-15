package pancake

import (
	"BWINF/utils/slice"
	"strconv"
	"strings"
)

// SortSteps is a slice of steps to sort a pancake
type SortSteps []int8

func ParseSortSteps(s string) SortSteps {
	var sortSteps SortSteps
	for _, line := range strings.Split(s, " ") {
		if line == "" {
			continue
		}
		step, _ := strconv.Atoi(line)
		sortSteps = append(sortSteps, int8(step))
	}
	return sortSteps
}

func (s SortSteps) String() string {
	stringSteps := slice.Map(s, func(e int8) string {
		return strconv.Itoa(int(e))
	})
	return strings.Join(stringSteps, " ")
}

func (s *SortSteps) Push(e int8) *SortSteps {
	*s = append(*s, e)
	return s
}

func (s *SortSteps) Copy() *SortSteps {
	newS := make(SortSteps, len(*s), cap(*s))
	copy(newS, *s)
	return &newS
}
