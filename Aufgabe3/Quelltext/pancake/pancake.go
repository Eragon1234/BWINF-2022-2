package pancake

import (
	"Aufgabe3/utils"
	"bufio"
	"io"
	"strconv"
)

type Pancake []int

func ParsePancakeFromReader(reader io.Reader) (Pancake, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	p := Pancake{}
	scanner.Scan() // ignoring first line because it represents the length which gets counted automatically
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return Pancake{}, err
		}
		p.Push(i)
	}
	utils.ReverseSlice(p) // reversing the whole stack because we parse it in reverse order
	return p, nil
}

// Flip flips the stack at the given index i and removes/eats the topmost pancake/element
func (p *Pancake) Flip(i int) {
	index := len(*p) - i
	utils.ReverseSlice((*p)[index:])
	_, *p = utils.Pop(*p) // removing/eating the topmost pancake
}

// Push adds an element to the stack and increases the length
func (p *Pancake) Push(e int) {
	*p = append(*p, e)
}
