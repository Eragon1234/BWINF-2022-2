package aufgabe3

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/cli"
	"bufio"
	"fmt"
	"os"
)

var Sort = cli.Command{
	Name:        "sort",
	Usage:       "sort <filename>",
	Description: "Berechnet den Sortierweg für den Stack in der Datei",
	Action:      sort,
}

func sort(args []string, c *cli.Command) error {
	file, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer file.Close()
	stack, err := pancake.ParseStack[int8](bufio.NewReader(file))
	if err != nil {
		return err
	}
	fmt.Printf("Stack: %v\n", stack)
	sortSteps := pancake.BruteForceSort(stack)
	for _, step := range sortSteps {
		fmt.Printf("flip bei %v, neuer Stack %v\n", step, stack.Flip(int(step)))
	}
	return nil
}
