package pancake

import (
	"Aufgabe3/utils"
	"bufio"
	"io"
	"strconv"
)

type SortSteps []int

type Stack []int

func Parse(reader io.Reader) (Stack, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	p := Stack{}
	scanner.Scan() // ignoring first line because it represents the length which gets counted automatically
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return Stack{}, err
		}
		p.Push(i)
	}
	utils.ReverseSlice(p) // reversing the whole stack because we parse it in reverse order
	return p, nil
}

// Flip flips the stack at the given index i and removes/eats the topmost pancake/element
func (p *Stack) Flip(i int) Stack {
	index := len(*p) - i
	utils.ReverseSlice((*p)[index:])
	_, *p = utils.Pop(*p) // removing/eating the topmost pancake
	return nil
}

// Push adds an element to the stack and increases the length
func (p *Stack) Push(e int) Stack {
	*p = append(*p, e)
	return nil
}

// Copy returns a copy of the pancake
func (p *Stack) Copy() Stack {
	newP := make(Stack, len(*p))
	copy(newP, *p)
	return newP
}
