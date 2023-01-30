package pancake

import (
	"Aufgabe3/utils"
	"Aufgabe3/utils/slice"
	"bufio"
	"io"
	"strconv"
)

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
	slice.ReverseSlice(p) // reversing the whole stack because we parse it in reverse order
	return p, nil
}

// Flip flips the stack at the given index i and removes/eats the topmost pancake/element
func (p *Stack[T]) Flip(i int) *Stack[T] {
	index := len(*p) - i
	slice.ReverseSlice((*p)[index:])
	_, *p = slice.Pop(*p) // removing/eating the topmost pancake
	return p
}

// Push adds an element to the stack and increases the length
func (p *Stack[T]) Push(e T) *Stack[T] {
	*p = append(*p, e)
	return p
}

// Copy returns a copy of the pancake
func (p *Stack[T]) Copy() *Stack[T] {
	newP := make(Stack[T], len(*p))
	copy(newP, *p)
	return &newP
}
