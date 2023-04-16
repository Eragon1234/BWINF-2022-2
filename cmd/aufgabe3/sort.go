package aufgabe3

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/Aufgabe3/pancake/sort"
	"BWINF/cli"
	"BWINF/pkg/set"
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

var Sort = cli.Command{
	Name:  "sort",
	Usage: "sort <filename>",
	Aliases: *set.FromSlice([]string{
		"sortPancake",
	}),
	Description: "Berechnet den Sortierweg für den Stack in der Datei",
	Action:      sortPancake,
}

func sortPancake(args []string, c *cli.Command) error {
	if len(args) == 0 {
		return errors.New("keine Datei angegeben")
	}
	filename := args[0]
	if !fs.ValidPath(filename) {
		return errors.New("ungültiger Dateipfad")
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	stack, err := pancake.ParseStack(bufio.NewReader(file))
	if err != nil {
		return err
	}
	fmt.Printf("Stack: %v\n", stack)
	sortSteps := sort.BruteForceMultiGoroutineInline(stack)
	for _, step := range sortSteps {
		fmt.Printf("flip bei %v, neuer Stack %v\n", step, stack.Flip(int(step)))
	}
	return nil
}
