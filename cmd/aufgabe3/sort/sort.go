package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/Aufgabe3/pancake/sort"
	"BWINF/cli"
	"BWINF/pkg/set"
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
	Subcommands: []cli.Command{
		Astar,
	},
}

func sortPancake(args []string, _ *cli.Command) error {
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	stack, err := pancake.ParseStack(file)
	if err != nil {
		return err
	}
	printStack(stack)
	sortSteps := sort.BruteForceInlined(stack)
	printSortSteps(stack, sortSteps)
	return nil
}

func printStack(stack pancake.Stack) {
	fmt.Printf("Stack: %v\n", stack)
}

func printSortSteps(stack pancake.Stack, sortSteps pancake.SortSteps) {
	for _, step := range sortSteps {
		fmt.Printf("flip bei %v, neuer Stack %v\n", step, stack.Flip(int(step)))
	}
}
