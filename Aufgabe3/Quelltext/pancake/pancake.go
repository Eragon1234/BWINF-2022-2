package pancake

import (
	"Aufgabe3/utils"
	"bufio"
	"io"
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

type Stack[T utils.Number] []T

func Parse[T utils.Number](reader io.Reader) (Stack[T], error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	p := Stack[T]{}
	scanner.Scan() // ignoring first line because it represents the length which gets counted automatically
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return Stack[T]{}, err
		}
		p.Push(T(i))
	}
	utils.ReverseSlice(p) // reversing the whole stack because we parse it in reverse order
	return p, nil
}

// Flip flips the stack at the given index i and removes/eats the topmost pancake/element
func (p *Stack[T]) Flip(i int) Stack[T] {
	index := len(*p) - int(i)
	utils.ReverseSlice((*p)[index:])
	_, *p = utils.Pop(*p) // removing/eating the topmost pancake
	return nil
}

// Push adds an element to the stack and increases the length
func (p *Stack[T]) Push(e T) Stack[T] {
	*p = append(*p, e)
	return nil
}

// Copy returns a copy of the pancake
func (p *Stack[T]) Copy() Stack[T] {
	newP := make(Stack[T], len(*p))
	copy(newP, *p)
	return newP
}
