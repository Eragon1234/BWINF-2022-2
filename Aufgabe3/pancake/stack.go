package pancake

import (
	"BWINF/pkg/slice"
	"bufio"
	"io"
	"strconv"
	"strings"
)

// KeepTrackOfSide is a global variable that is used to control if the side of the pancakes should be kept track of.
var KeepTrackOfSide = false

// Stack represents a stack of pancakes/elements
type Stack []int8

func ParseStack(reader io.Reader) (Stack, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	p := Stack{}
	scanner.Scan() // ignoring the first line because it represents the length which gets counted automatically
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return Stack{}, err
		}
		p.Push(int8(i))
	}
	slice.ReverseSlice(p) // reversing the whole stack because we parse it in reverse order
	return p, nil
}

func (s Stack) String() string {
	stringSteps := slice.Map(s, func(e int8) string {
		return strconv.Itoa(int(e))
	})
	return strings.Join(stringSteps, " ")
}

// Flip flips the stack at the given index i and removes/eats the topmost pancake/element
func (s *Stack) Flip(i int) *Stack {
	index := len(*s) - i
	slice.ReverseSlice((*s)[index:])
	_, *s = slice.Pop(*s) // removing/eating the topmost pancake
	if KeepTrackOfSide {
		// flip the signs of the reversed part
		for i := index; i < len(*s); i++ {
			(*s)[i] = -(*s)[i]
		}
	}
	return s
}

// Push adds an element to the stack and increases the length
func (s *Stack) Push(e int8) *Stack {
	*s = append(*s, e)
	return s
}

// Copy returns a copy of the pancake
func (s *Stack) Copy() *Stack {
	newP := make(Stack, len(*s))
	copy(newP, *s)
	return &newP
}
