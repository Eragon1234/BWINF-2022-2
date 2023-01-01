package pancake

import (
	"Aufgabe3/utils"
	"bufio"
	"os"
	"strconv"
)

type Pancake struct {
	Stack  []int
	Length int
}

func ParsePancakeFromFile(filename string) (Pancake, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Pancake{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	p := Pancake{}
	scanner.Scan() // ignoring first line because it represents the length
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return Pancake{}, err
		}
		p.Push(i)
	}
	return p, nil
}

func (p *Pancake) Flip(i int) {
	index := len(p.Stack) - i
	utils.ReverseSlice(p.Stack[index:])
}

func (p *Pancake) Push(e int) {
	p.Stack = append(p.Stack, e)
	p.Length++
}
